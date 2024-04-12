package service

import (
	"encoding/json"
	"file-service/common"
	"file-service/repository"
	"file-service/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	FileUploading     = "文件上传中"
	FileUploadFail    = "文件上传失败,失败原因 "
	FileUploadSuccess = "文件上传成功"
)

type IFileUploadService interface {
	FileUpload(ctx *gin.Context, header *multipart.FileHeader, lock string)
}

// FileUploadService 定义空结构体
type FileUploadService struct {
	FileRepository repository.IFileRepository
}

// NewFileUploadService 构造函数·
func NewFileUploadService() IFileUploadService {
	return FileUploadService{
		repository.NewFileRepository(),
	}
}

// FileUpload 文件上传方法
func (f FileUploadService) FileUpload(ctx *gin.Context, file *multipart.FileHeader, lock string) {
	//获取单机锁并设置失效时间
	locked := utils.GetLockAndExpire(lock, time.Second*5)
	if !locked {
		ctx.JSON(http.StatusOK, common.OfFail(common.UploadedFail))
		return
	}
	fileModel := f.FileRepository.Create(file.Filename, 1)
	// 读取文件
	open, err := file.Open()
	if err != nil {
		log.Println("Open File Fail", err)
		ctx.JSON(http.StatusInternalServerError, common.OfFail(common.SystemFail))
		return
	}
	//回调
	ctx.JSON(http.StatusOK, common.OfSuccess(FileUploading))
	//异步
	go func() {
		data := map[string]any{
			"FileId": fileModel.Id,
			"UserId": fileModel.UserId,
		}
		//文件to字节数组
		bytes, err := io.ReadAll(open)
		//上传字节数组to七牛云
		hash, err := utils.OssFileUpload(bytes)
		if err != nil {
			log.Println("FileUpload Upload Oss File", err)
			fileModel.Status = 2
			f.FileRepository.Updates(fileModel)
			data["Message"] = FileUploadFail + err.Error()
			marshal, _ := json.Marshal(data)
			utils.SendSync(common.TopicFileUploadNotice, marshal)
		} else {
			log.Println("FileUpload Upload Oss Success:")
			//更新为上传成功的状态
			fileModel.Status = 3
			fileModel.Hash = hash
			f.FileRepository.Updates(fileModel)
			data["Message"] = FileUploadSuccess
			marshal, _ := json.Marshal(data)
			utils.SendSync(common.TopicFileUploadNotice, marshal)
		}
	}()
}

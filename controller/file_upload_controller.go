package controller

import (
	"encoding/json"
	"file-service/common"
	"file-service/repository"
	"file-service/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type IFileUploadController interface {
	FileUpload(context *gin.Context)
}

type FileUploadController struct {
	UserRepository repository.IUserRepository
	FileRepository repository.IFileRepository
}

func NewFileUploadController() IFileUploadController {
	return FileUploadController{
		repository.NewUserRepository(),
		repository.NewFileRepository(),
	}
}

func (f FileUploadController) FileUpload(context *gin.Context) {
	systemFail := common.Fail{
		Code: 10001,
		Msg:  "系统繁忙",
	}
	file, err := context.FormFile("file")
	if err != nil {
		log.Println("FileUpload Error:", err)
		context.JSON(http.StatusInternalServerError, common.OfFail(systemFail))
		return
	}
	if file.Size > 8<<20 {
		context.JSON(http.StatusBadRequest, common.OfFail(common.Fail{Code: 10002, Msg: "文件过大"}))
		return
	}
	log.Println("FileUpload req:", file.Filename)
	fileModel := f.FileRepository.Create(file.Filename, 1)
	open, err := file.Open()
	if err != nil {
		log.Println("Open File Fail", err)
		context.JSON(http.StatusInternalServerError, common.OfFail(systemFail))
		//TODO return和context.Next()效果一样？
		return
	}
	context.JSON(http.StatusOK, common.OfSuccess("上传中"))
	go func() {
		data := map[string]any{
			"FileId": fileModel.Id,
			"UserId": fileModel.UserId,
		}
		all, err := io.ReadAll(open)
		hash, err := utils.OssFileUpload(all)
		if err != nil {
			log.Println("Upload Oss File", err)
			fileModel.Status = 2
			f.FileRepository.Updates(fileModel)
			data["Message"] = "上传失败,失败原因:" + err.Error()
			marshal, err := json.Marshal(data)
			if err != nil {
				return
			}
			utils.SendSync(common.TopicFileUploadNotice, marshal)
			return
		}
		//更新为上传成功的状态
		fileModel.Status = 3
		fileModel.Hash = hash
		f.FileRepository.Updates(fileModel)
		data["Message"] = "上传成功"
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		utils.SendSync(common.TopicFileUploadNotice, marshal)
	}()
}

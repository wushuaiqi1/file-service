package controller

import (
	"encoding/json"
	"file-service/common"
	"file-service/repository"
	"file-service/service"
	"file-service/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type IFileUploadController interface {
	FileUpload(ctx *gin.Context)
	FileUploadLock(ctx *gin.Context)
}

// FileUploadController 引入对应接口，而不是具体实现
type FileUploadController struct {
	FileRepository    repository.IFileRepository
	FileUploadService service.IFileUploadService
}

func NewFileUploadController() IFileUploadController {
	return FileUploadController{
		repository.NewFileRepository(),
		service.NewFileUploadService(),
	}
}

func (f FileUploadController) FileUploadLock(ctx *gin.Context) {
	userId, file, ok := fileUploadReqParamsCheck(ctx)
	if !ok {
		return
	}
	f.FileUploadService.FileUpload(ctx, file, userId)
}

func (f FileUploadController) FileUpload(ctx *gin.Context) {
	_, file, ok := fileUploadReqParamsCheck(ctx)
	if !ok {
		return
	}
	fileModel := f.FileRepository.Create(file.Filename, 1)
	open, err := file.Open()
	if err != nil {
		log.Println("Open File Fail", err)
		ctx.JSON(http.StatusInternalServerError, common.OfFail(common.SystemFail))
		return
	}
	ctx.JSON(http.StatusOK, common.OfSuccess("上传中"))
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

// FileUploadReqParamsCheck 文件上传请求前置校验
func fileUploadReqParamsCheck(ctx *gin.Context) (userId string, file *multipart.FileHeader, ok bool) {
	userId = ctx.PostForm("userId")
	file, err := ctx.FormFile("file")
	if userId == "" || err != nil {
		ctx.JSON(http.StatusOK, common.OfFail(common.MissingParam))
		return "", nil, ok
	}
	if file.Size > 8<<20 {
		ctx.JSON(http.StatusBadRequest, common.OfFail(common.BodySizeLimit))
		return "", nil, ok
	}
	return userId, file, true
}

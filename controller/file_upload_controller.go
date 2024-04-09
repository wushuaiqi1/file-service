package controller

import (
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
	log.Println("FileUpload req:", file.Filename)
	fileModel := f.FileRepository.Create(file.Filename, 1)
	open, err := file.Open()
	if err != nil {
		log.Println("Open File Fail", err)
		context.JSON(http.StatusInternalServerError, common.OfFail(systemFail))
		//TODO return和context.Next()效果一样？
		return
	}
	all, err := io.ReadAll(open)
	hash, err := utils.OssFileUpload(all)
	if err != nil {
		log.Println("Upload Oss File", err)
		fileModel.Status = 2
		f.FileRepository.Updates(fileModel)
		return
	}
	fileModel.Status = 3
	fileModel.Hash = hash
	f.FileRepository.Updates(fileModel)
	context.JSON(http.StatusOK, common.OfSuccess(nil))
}

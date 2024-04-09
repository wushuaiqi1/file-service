package repository

import (
	"file-service/common"
	"file-service/model"
)

type IFileRepository interface {
	Updates(file *model.File)
	Create(name string, userId uint64) (file *model.File)
}

type FileRepository struct {
}

func (f FileRepository) Updates(file *model.File) {
	common.DbInstance.Updates(file)
}

func (f FileRepository) Create(name string, userId uint64) (file *model.File) {
	file = &model.File{
		Name:   name,
		UserId: userId,
		Status: 1, //默认状态是上传中
	}
	common.DbInstance.Create(file)
	return file
}

func NewFileRepository() IFileRepository {
	return FileRepository{}
}

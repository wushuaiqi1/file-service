package repository

import (
	"log"
	"notice-service/common"
	"notice-service/models"
)

type INoticeRepository interface {
	Create(notice *models.Notice)
}

type NoticeRepository struct {
}

func NewNoticeRepository() INoticeRepository {
	return NoticeRepository{}
}

func (nr NoticeRepository) Create(notice *models.Notice) {
	common.DbInstance.Create(notice)
	err := common.DbInstance.Error
	if err != nil {
		log.Println("Create Error:", err)
	}
}

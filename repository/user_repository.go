package repository

import (
	"file-service/common"
	"file-service/model"
	"log"
)

type IUserRepository interface {
	GetUserById(id int64) (user *model.User)
}

type UserRepository struct {
}

func (u UserRepository) GetUserById(id int64) (user *model.User) {
	res := common.DbInstance.Where("id=?", id).First(&user)
	if res.Error != nil {
		log.Println("GetUserById Error ", res.Error)
		return nil
	}
	return user
}

func NewUserRepository() IUserRepository {
	return UserRepository{}
}

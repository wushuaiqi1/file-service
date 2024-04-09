package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func InitDatabase() error {
	dsn := "root:@tcp(127.0.0.1:3306)/forge?charset=utf8mb4&parseTime=False&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DbInstance = db
	return nil
}

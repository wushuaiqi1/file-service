package model

type User struct {
	Id        uint64 `gorm:"primarykey"`
	CreatedAt uint64
	UpdatedAt uint64
	IsDeleted uint8
	UserName  string
	Password  string
}

func (u User) TableName() string {
	return "user"
}

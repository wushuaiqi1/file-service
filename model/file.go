package model

type File struct {
	Id        uint64 `gorm:"primarykey"`
	CreatedAt uint64
	UpdatedAt uint64
	IsDeleted uint8
	Hash      string
	Name      string
	Status    uint8
	UserId    uint64
}

type Tabler interface {
	TableName() string
}

func (f File) TableName() string {
	return "file"
}

package models

type Notice struct {
	Id        uint64
	Type      uint8
	Message   string
	UserId    uint64
	FileId    uint64
	CreatedAt uint64
	UpdatedAt uint64
	IsDeleted uint8
}

type Tabler interface {
	TableName() string
}

func (n Notice) TableName() string {
	return "notice"
}

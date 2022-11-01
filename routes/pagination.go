package routes

import (
	"gorm.io/gorm"
)

type PaginatedResponse struct {
	Page      uint
	PageSize  uint
	TotalSize uint
	LastPage  bool
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

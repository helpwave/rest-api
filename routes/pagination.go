package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type PaginatedResponse struct {
	Page      int
	PageSize  int
	TotalSize int64
	LastPage  bool
}

func Paginate(pagination PaginatedResponse) func(db *gorm.DB) *gorm.DB {
	page := pagination.Page
	pageSize := pagination.PageSize

	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetPagination(ctx *gin.Context, db *gorm.DB, model interface{}) (PaginatedResponse, error) {
	q := ctx.Request.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize > 100 {
		pageSize = 100
	}

	var totalSize int64
	tx := db.Model(model).Count(&totalSize)

	if tx.Error != nil {
		return PaginatedResponse{}, tx.Error
	}

	lastPage := int64(page*pageSize) >= totalSize

	return PaginatedResponse{
		Page:      page,
		PageSize:  pageSize,
		TotalSize: totalSize,
		LastPage:  lastPage,
	}, nil

}

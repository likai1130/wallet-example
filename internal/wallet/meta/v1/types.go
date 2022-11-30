package v1

import (
	"time"
)

type Extend map[string]interface{}

const defaultPageSize = 10
const defaultPageIndex = 1

type ListMeta struct {
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	TotalCount  int64 `json:"total_count"`
}

func (meta *ListMeta) GetCurrentPage() int64 {
	return meta.CurrentPage
}

func (meta *ListMeta) SetCurrentPage(offset int64) {
	meta.CurrentPage = offset
}

func (meta *ListMeta) GetTotalPage() int64 {
	return meta.TotalCount
}

func (meta *ListMeta) SetTotalPage(limit int64) {
	totalCount := meta.TotalCount
	totalPages := int64(0)
	if meta.TotalCount > 0 {
		if totalCount%limit != 0 {
			totalPages = (totalCount / limit) + 1
		} else {
			totalPages = totalCount / limit
		}
	}
	meta.TotalPage = totalPages
}

type Model struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time `json:"createdAt" gorm:"index;comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index;comment:最后更新时间"`
}

type ObjectMeta struct {
	ID     uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Extend Extend `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ModelTime
}

type ListOptions struct {
	PageIndex int64 `form:"pageIndex"`
	PageSize  int64 `form:"pageSize"`
}

func (l *ListOptions) GetPageSize() int64 {
	if l.PageSize <= 0 {
		l.PageSize = defaultPageSize
	}
	return l.PageSize
}

func (l *ListOptions) GetPageIndex() int64 {
	if l.PageIndex <= 0 {
		l.PageIndex = defaultPageIndex
	}
	offset := (l.PageIndex - 1) * l.PageSize
	if offset < 0 {
		offset = 0
	}
	return offset
}

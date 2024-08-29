package pagination

import (
	"math"

	"github.com/MQEnergy/go-skeleton/internal/vars"
)

// PaginateResp ...
type PaginateResp struct {
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
	List        any   `json:"list"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}

func New() *PaginateResp {
	return &PaginateResp{}
}

// ParsePage 分页超限设置和格式化
func (pb *PaginateResp) ParsePage(currentPage, pageSize int) (page PaginateResp) {
	// 返回每页数量
	page.PerPage = pageSize
	// 返回当前页码
	page.CurrentPage = currentPage

	if currentPage < 1 {
		page.CurrentPage = 1
	}
	if pageSize < 1 {
		page.PerPage = vars.Config.GetInt("server.defaultPageSize")
	}
	if pageSize > vars.Config.GetInt("server.maxPageSize") {
		page.PerPage = vars.Config.GetInt("server.maxPageSize")
	}
	if page.LastPage < 1 {
		page.LastPage = 1
	}
	return
}

// GetOffset 获取偏移量
func (pb *PaginateResp) GetOffset() int {
	return (pb.CurrentPage - 1) * pb.PerPage
}

// GetLimit 获取每页数量
func (pb *PaginateResp) GetLimit() int {
	return pb.PerPage
}

// GetLastPage 计算总页数
func (pb *PaginateResp) GetLastPage() int {
	if pb.Total > int64(pb.PerPage) {
		pb.LastPage = int(math.Ceil(float64(pb.Total) / float64(pb.PerPage)))
	}
	return pb.LastPage
}

// SetList 设置列表数据
func (pb *PaginateResp) SetList(list any) *PaginateResp {
	pb.List = list
	return pb
}

// SetCount 设置总数
func (pb *PaginateResp) SetCount(count int64) *PaginateResp {
	pb.Total = count
	return pb
}

package view

import (
	"github.com/aganomaminmi/go-mvc/pkg/model"
)

type PageInfoView struct {
	Count       int64 `json:"count"`
	PageSize    int64 `json:"pageSize"`
	CurrentPage int64 `json:"currentPage"`
	TotalPage   int64 `json:"totalPage"`
}

func NewPageInfoView(p model.Page) PageInfoView {
	return PageInfoView{
		Count:       p.Count,
		PageSize:    p.PageSize,
		CurrentPage: p.CurrentPage,
		TotalPage:   p.TotalPage,
	}
}

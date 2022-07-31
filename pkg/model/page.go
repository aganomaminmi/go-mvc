package model

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type Page struct {
	Count       int64
	TotalPage   int64
	CurrentPage int64
	PageSize    int64
}

func (p *Page) CalcTotalPage(rws int64) {
	p.Count = rws
	p.TotalPage = int64(math.Ceil(float64(rws) / float64(p.PageSize)))

}

func (p Page) Pagenate() (func(db *gorm.DB) *gorm.DB, error) {
	scps := func(db *gorm.DB) *gorm.DB {

		offset := (int(p.CurrentPage) - 1) * int(p.PageSize)
		return db.Offset(offset).Limit(int(p.PageSize))
	}

	if p.PageSize > 100 {
		return scps, fmt.Errorf("page_size has been under 101.")
	}
	if p.PageSize <= 0 {
		return scps, fmt.Errorf("page_size has been over 0.")
	}

	return scps, nil
}

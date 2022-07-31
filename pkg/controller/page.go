package controller

import (
	"net/http"
	"strconv"
)

type PagingQuery struct {
	Page     int64
	PageSize int64
}

func NewPagingQuery(r *http.Request) (PagingQuery, error) {
	q := r.URL.Query()
	pg, err := strconv.Atoi(q.Get("page"))
	if err != nil {
		return PagingQuery{}, err
	}
	if pg == 0 {
		pg = 1
	}

	pgSz, err := strconv.Atoi(q.Get("page_size"))
	if err != nil {
		return PagingQuery{}, err
	}

	return PagingQuery{
		Page:     int64(pg),
		PageSize: int64(pgSz),
	}, nil

}

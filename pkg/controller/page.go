package controller

import (
	"net/http"
	"strconv"
)

type PagingQuery struct {
	Page     int64
	PageSize int64
}

func NewPagingQuery(r *http.Request) PagingQuery {
	q := r.URL.Query()
	pg, err := strconv.Atoi(q.Get("page"))
	if err != nil {
		pg = 1
	}
	if pg == 0 {
		pg = 1
	}

	pgSz, err := strconv.Atoi(q.Get("page_size"))
	if err != nil {
		pgSz = 10
	}

	return PagingQuery{
		Page:     int64(pg),
		PageSize: int64(pgSz),
	}

}

package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Body          func(obj any) error
	GetParam      func(key string) string
	GetQueryParam func(key string) string
	GetHeader     func(key string) string
}

func NewRequest(ctx *gin.Context) Request {
	return Request{
		Body:          ctx.BindJSON,
		GetParam:      ctx.Param,
		GetQueryParam: ctx.Query,
		GetHeader:     ctx.Request.Header.Get,
	}
}

type Paginate struct {
	page     int
	pageSize int
}

func (p Paginate) GetOffset() int {
	return (p.page - 1) * p.pageSize
}

func NewPaginate(req Request) Paginate {
	page, _ := strconv.Atoi(req.GetQueryParam("page"))
	pageSize, _ := strconv.Atoi(req.GetQueryParam("page_size"))

	return Paginate{
		page,
		pageSize,
	}
}

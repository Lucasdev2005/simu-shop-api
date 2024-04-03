package core

import "github.com/gin-gonic/gin"

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

func (r Request) NewFail(errorCode int, message string) Fail {
	return Fail{
		ErrorCode: errorCode,
		Response:  message,
	}
}

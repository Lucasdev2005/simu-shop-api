package handler

import (
	"net/http"
	"simushop/internal/core"
)

func (h handler) CreateUser(request core.Request) (core.Success, core.Fail) {
	return core.Success{}, core.Fail{
		ErrorCode: http.StatusNotImplemented,
		Message:   "not implemented",
	}
}

package handler

import (
	"simushop/internal/core"
)

type CreateUser struct {
}

func (h handler) CreateUser(request core.Request) (core.Success, core.Fail) {
	var user CreateUser
	request.Body(&user)

	return core.Success{
		SuccessCode: 200,
		Data:        user,
	}, core.Fail{}
}

package handler

import (
	"simushop/internal/apirest"
	"simushop/internal/core"
	"simushop/internal/entity"
)

type CreateUser struct {
	Username     string  `json:"username"`
	UserBalance  float64 `json:"user_balance"`
	UserPassword string  `json:"user_password"`
	UserType     string  `json:"user_type"`
}

func (h handler) CreateUser(request core.Request) (core.Success, core.Fail) {
	var user CreateUser
	request.Body(&user)

	err := h.repository.CreateUser(entity.User{
		UserPassword: user.UserPassword,
		Username:     user.Username,
		UserBalance:  user.UserBalance,
		UserType:     user.UserType,
	})

	if err != nil {
		return core.Success{}, apirest.BadRequest(err.Error())
	}

	return apirest.Created(user), core.Fail{}

}

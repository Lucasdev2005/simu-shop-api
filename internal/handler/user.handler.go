package handler

import (
	"simushop/internal/core"
	"simushop/internal/entity"
)

type CreateUserDTO struct {
	Username     string  `json:"username"`
	UserBalance  float64 `json:"user_balance"`
	UserPassword string  `json:"user_password"`
	UserType     string  `json:"user_type"`
}

func (h handler) CreateUser(request core.Request) core.Response {
	var user CreateUserDTO
	request.Body(&user)

	err := h.repository.CreateUser(entity.User{
		UserPassword: user.UserPassword,
		Username:     user.Username,
		UserBalance:  user.UserBalance,
		UserType:     user.UserType,
	})

	if err != nil {
		return core.BadRequest(err.Error())
	}

	return core.Created(user)

}

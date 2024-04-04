package handler

import (
	"simushop/internal/core"
	"simushop/internal/entity"
)

type CreateUserDTO struct {
	Username     string  `json:"username" validate:"required"`
	UserBalance  float64 `json:"user_balance" validate:"required"`
	UserPassword string  `json:"user_password" validate:"required"`
	UserType     string  `json:"user_type" validate:"required,userType"`
}

func (c CreateUserDTO) ToEntity() entity.User {
	return entity.User{
		UserPassword: c.UserPassword,
		Username:     c.Username,
		UserBalance:  c.UserBalance,
		UserType:     c.UserType,
	}
}

func (h handler) CreateUser(request core.Request) core.Response {
	var user CreateUserDTO
	request.Body(&user)
	errFromDTO := h.ValidateStruct(user)

	if errFromDTO != nil {
		return core.BadRequest(errFromDTO.Error())
	}

	err := h.repository.CreateUser(user.ToEntity())

	if err != nil {
		return core.BadRequest(err.Error())
	}

	return core.Created(user)
}

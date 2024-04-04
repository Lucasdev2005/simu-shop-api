package dto

import "simushop/internal/entity"

type CreateUserDTO struct {
	Username     string  `json:"username" validate:"required"`
	UserBalance  float64 `json:"user_balance" validate:"required,min=0"`
	UserPassword string  `json:"user_password" validate:"required"`
	UserType     string  `json:"user_type" validate:"required,userType"`
}

func (c CreateUserDTO) ToUser() entity.User {
	return entity.User{
		UserPassword: c.UserPassword,
		Username:     c.Username,
		UserBalance:  c.UserBalance,
		UserType:     c.UserType,
	}
}

package dto

import "simushop/internal/entity"

type CreateUserDTO struct {
	Username     string  `json:"username" validate:"required"`
	UserBalance  float64 `json:"userBalance" validate:"required,min=0"`
	UserPassword string  `json:"userPassword" validate:"required"`
	UserType     string  `json:"userType" validate:"required,userType"`
}

type UpdateUserDTO struct {
	Username     string  `json:"username"`
	UserBalance  float64 `json:"userBalance" validate:"min=0"`
	UserPassword string  `json:"userPassword"`
	UserType     string  `json:"userType,omitempty" validate:"updateUserType"`
}

func (c CreateUserDTO) ToUser() entity.User {
	return entity.User{
		UserPassword: c.UserPassword,
		Username:     c.Username,
		UserBalance:  c.UserBalance,
		UserType:     c.UserType,
	}
}

func (c UpdateUserDTO) ToUser() entity.User {
	return entity.User{
		UserPassword: c.UserPassword,
		Username:     c.Username,
		UserBalance:  c.UserBalance,
		UserType:     c.UserType,
	}
}

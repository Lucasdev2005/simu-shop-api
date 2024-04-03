package handler

import (
	"simushop/internal/entity"

	"github.com/go-playground/validator/v10"
)

type repository interface {
	CreateUser(user entity.User) error
}

type handler struct {
	repository repository
	validate   *validator.Validate
}

func NewHandler(repository repository) handler {
	return handler{
		repository,
		validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h handler) ValidateStruct(s interface{}) error {
	return h.validate.Struct(s)
}

package handler

import (
	"simushop/internal/entity"
	"simushop/internal/handler/validators"

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
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation(validators.ValidUserType())
	return handler{
		repository,
		validate,
	}
}

func (h handler) ValidateStruct(s interface{}) error {
	return h.validate.Struct(s)
}

package handler

import (
	"simushop/internal/entity"
	"simushop/internal/handler/validators"

	"github.com/go-playground/validator/v10"
)

type repository interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User, where string, args ...interface{}) error

	CreateProduct(p entity.Product) error
	UpdateProduct(p entity.Product, where string, args ...interface{}) error

	CreateTopic(topic entity.Topic) error
	UpdateTopic(topic entity.Topic, where string, args ...interface{}) error
	ListTopics(where string, args ...interface{}) []entity.Topic
}

type handler struct {
	repository repository
	validate   *validator.Validate
}

func NewHandler(repository repository) handler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation(validators.ValidUserType())
	validate.RegisterValidation(validators.UptateValidUserType())
	return handler{
		repository,
		validate,
	}
}

func (h handler) ValidateStruct(s interface{}) error {
	return h.validate.Struct(s)
}

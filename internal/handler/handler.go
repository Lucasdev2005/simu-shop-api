package handler

import "simushop/internal/entity"

type repository interface {
	CreateUser(user entity.User)
}

type handler struct {
	repository repository
}

func NewHandler(repository repository) handler {
	return handler{repository}
}

package handler

type repository interface {
	CreateUser()
}

type handler struct {
	repository repository
}

func NewHandler(repository repository) handler {
	return handler{repository}
}

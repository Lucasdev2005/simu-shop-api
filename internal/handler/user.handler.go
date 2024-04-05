package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
)

func (h handler) CreateUser(request core.Request) core.Response {
	var user dto.CreateUserDTO
	request.Body(&user)
	errFromDTO := h.ValidateStruct(user)

	if errFromDTO != nil {
		return core.BadRequest(errFromDTO.Error())
	}

	err := h.repository.CreateUser(user.ToUser())

	if err != nil {
		return core.BadRequest(err.Error())
	}

	return core.Created(user)
}

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

	userCreated, err := h.repository.CreateUser(user.ToUser())

	if err != nil {
		return core.BadRequest(err.Error())
	}

	return core.Created(userCreated)
}

func (h handler) UpdateUser(request core.Request) core.Response {
	var user dto.UpdateUserDTO
	request.Body(&user)

	errFromDTO := h.ValidateStruct(user)

	if errFromDTO != nil {
		return core.BadRequest(errFromDTO.Error())
	}

	err := h.repository.UpdateUser(user.ToUser(), "user_id = ?", request.GetParam("id"))

	if err != nil {
		return core.BadRequest(err.Error())
	}

	return core.Ok("user updated")
}

package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
)

func (h handler) AddItemOnCart(request core.Request) core.Response {
	var (
		body dto.AddItemCartDTO
	)

	request.Body(&body)
	err := h.ValidateStruct(body)

	if err != nil {
		return core.BadRequest(err.Error())
	}

	shoppingCart, err := h.repository.AddItemOnCart(body.ToShoppingCart())

	if err != nil {
		return core.InternalError(err.Error())
	} else {
		return core.Created(shoppingCart)
	}
}

func (h handler) RemoveItemFromCart(request core.Request) core.Response {
	return core.Ok("")
}

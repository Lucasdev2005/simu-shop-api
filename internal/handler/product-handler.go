package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
)

func (h handler) CreateProduct(request core.Request) core.Response {
	var body dto.CreateProductDTO
	request.Body(&body)
	errOnDTO := h.ValidateStruct(body)

	if errOnDTO != nil {
		return core.BadRequest(errOnDTO)
	}

	return core.Ok("")
}

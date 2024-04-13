package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
)

func (h handler) CreateProduct(request core.Request) core.Response {
	var (
		body   dto.CreateProductDTO
		result core.Response = core.Ok("Product Created.")
	)
	request.Body(&body)
	errOnDTO := h.ValidateStruct(body)

	if errOnDTO != nil {
		result = core.BadRequest(errOnDTO)
	} else {
		errOnCreateProduct := h.repository.CreateProduct(body.ToProduct())
		if errOnCreateProduct != nil {
			result = core.BadRequest(errOnCreateProduct.Error())
		}
	}

	return result
}

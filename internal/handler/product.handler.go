package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"strconv"
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

func (h handler) UpdateProduct(request core.Request) core.Response {
	var (
		body   dto.UpdateProductDTO
		result core.Response = core.Ok("Product Updated.")
	)
	request.Body(&body)

	errOnDTO := h.ValidateStruct(body)

	parsedId, _ := strconv.Atoi(request.GetParam("id"))

	if errOnDTO != nil {
		result = core.BadRequest(errOnDTO.Error())
	} else {
		errOnUpdate := h.repository.UpdateProduct(body.ToProduct(parsedId), "product_id = ?", request.GetParam("id"))
		if errOnUpdate != nil {
			result = core.BadRequest(errOnUpdate)
		}
	}

	return result
}

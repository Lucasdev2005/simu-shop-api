package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"simushop/internal/entity"
	"strconv"
)

func (h handler) CreateProduct(request core.Request) core.Response {
	var (
		body dto.CreateProductDTO
	)
	request.Body(&body)
	errOnDTO := h.ValidateStruct(body)

	if errOnDTO != nil {
		return core.BadRequest(errOnDTO)
	} else {
		product, errOnCreateProduct := h.repository.CreateProduct(body.ToProduct())
		if errOnCreateProduct != nil {
			return core.BadRequest(errOnCreateProduct.Error())
		} else {
			return core.Created(product)
		}
	}

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

func (h handler) ListProducts(request core.Request) core.Response {

	var (
		paginate    = core.NewPaginate(request)
		productname = request.GetQueryParam("productName")
		products    []entity.Product
	)

	if len(productname) > 0 {
		products = h.repository.PaginateTopics(paginate, "product_name LIKE ?", "%"+productname+"%")
	} else {
		products = h.repository.PaginateTopics(paginate, "")
	}

	return core.Ok(products)
}

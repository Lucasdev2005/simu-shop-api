package handler_test

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"testing"
)

func TestRequiredFields(t *testing.T) {
	t.Log("testing required fields.")

	res := createProduct(dto.CreateProductDTO{})

	BadRequest(res, t)
}

func TestCreateProduct(t *testing.T) {
	t.Log("testing creating a product.")

	res := createProduct(dto.CreateProductDTO{
		ProductName:             "Produto",
		ProductValue:            123,
		ProductCentimeterWidth:  123,
		ProductDescription:      "Descrição de teste",
		ProductKgWeitght:        123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 123,
		ProductDiscountPercent:  1,
		ProductSellerId:         1,
		TopicsIds:               []int{1, 2, 3},
	})

	Ok(res, t)
}

func TestUpdateProduct(t *testing.T) {
	t.Log("testing update product.")

	res := updateProduct(dto.UpdateProductDTO{
		ProductName: "Lucas Moreira exclusive",
	})

	Ok(res, t)
}

func TestMinValuesFromUpdate(t *testing.T) {
	t.Log("testing min values from update Product.")

	res := updateProduct(dto.UpdateProductDTO{
		ProductDescription: "aa",
	})

	BadRequest(res, t)
}

func createProduct(productDTO dto.CreateProductDTO) core.Response {
	return handlerInstance.CreateProduct(core.Request{
		Body: func(obj any) error {
			createProductDTO, ok := obj.(*dto.CreateProductDTO)

			if ok {
				createProductDTO.ProductName = productDTO.ProductName
				createProductDTO.ProductCentimeterHeight = productDTO.ProductCentimeterHeight
				createProductDTO.ProductCentimeterLength = productDTO.ProductCentimeterLength
				createProductDTO.ProductCentimeterWidth = productDTO.ProductCentimeterLength
				createProductDTO.ProductDescription = productDTO.ProductDescription
				createProductDTO.ProductValue = productDTO.ProductValue
				createProductDTO.ProductKgWeitght = productDTO.ProductKgWeitght
				createProductDTO.ProductSellerId = productDTO.ProductSellerId
				createProductDTO.TopicsIds = productDTO.TopicsIds
			}

			return nil
		},
	})
}

func updateProduct(data dto.UpdateProductDTO) core.Response {
	return handlerInstance.UpdateProduct(core.Request{
		Body: func(obj any) error {
			updateProductDTO, ok := obj.(*dto.UpdateProductDTO)

			if ok {
				updateProductDTO.ProductName = data.ProductName
				updateProductDTO.ProductCentimeterHeight = data.ProductCentimeterHeight
				updateProductDTO.ProductCentimeterLength = data.ProductCentimeterLength
				updateProductDTO.ProductCentimeterWidth = data.ProductCentimeterLength
				updateProductDTO.ProductDescription = data.ProductDescription
				updateProductDTO.ProductValue = data.ProductValue
				updateProductDTO.ProductKgWeitght = data.ProductKgWeitght
				updateProductDTO.ProductSellerId = data.ProductSellerId
			}

			return nil
		},
		GetParam: func(key string) string {
			return "1"
		},
	})
}

func TestGetTopics(t *testing.T) {
	t.Log("testing get Topics")
	res := handlerInstance.ListProducts(core.Request{
		GetQueryParam: func(key string) string {
			return "1"
		},
	})

	Ok(res, t)
}

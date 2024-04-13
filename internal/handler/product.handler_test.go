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
	})

	Ok(res, t)
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
			}

			return nil
		},
	})

}

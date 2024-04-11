package handler_test

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"testing"
)

func TestRequiredFields(t *testing.T) {
	t.Log("testing required fields.")

	res := handlerInstance.CreateProduct(core.Request{
		Body: func(obj any) error {
			createProductDTO, ok := obj.(*dto.CreateProductDTO)

			if ok {
				createProductDTO.ProductName = "produto de teste"
				createProductDTO.ProductCentimeterHeight = 2.6
				createProductDTO.ProductCentimeterLength = 2.6
				createProductDTO.ProductCentimeterWidth = 2.6
				createProductDTO.ProductDescription = "Produto de teste"
				createProductDTO.ProductValue = 2000.5
				createProductDTO.ProductKgWeitght = 456.2
			}

			return nil
		},
	})

	Ok(res, t)
}

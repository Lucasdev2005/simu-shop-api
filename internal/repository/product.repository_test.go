package repository_test

import (
	"simushop/internal/entity"
	"testing"
)

func TestCreatingProduct(t *testing.T) {
	t.Log("Creating Product")

	err := mockRepository.CreateProduct(entity.Product{
		ProductName:             "Produto de Teste",
		ProductValue:            123,
		ProductDiscountPercent:  2,
		ProductDescription:      "Descrição Fodastica",
		ProductKgWeitght:        123,
		ProductCentimeterWidth:  123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 35,
	})

	equal(t, nil, err, "Error on saving user on Database.")
}

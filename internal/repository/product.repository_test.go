package repository_test

import (
	"simushop/internal/entity"
	"testing"
)

func TestCreatingProduct(t *testing.T) {
	t.Log("Creating Product")

	err := createProduct()

	equal(t, nil, err, "Error on saving user on Database.")
}

func TestTryCreatingProductsWithsSameName(t *testing.T) {
	t.Log("testing creating products with same names.")

	/* creating product*/
	createProduct()

	/* creating product with same name*/
	err := createProduct()

	notEqual(t, nil, err, "Saving user with same names into the database.")
}

func createProduct() error {
	return mockRepository.CreateProduct(entity.Product{
		ProductName:             "Produto de Teste",
		ProductValue:            123,
		ProductDiscountPercent:  2,
		ProductDescription:      "Descrição de teste",
		ProductKgWeitght:        123,
		ProductCentimeterWidth:  123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 35,
	})
}

package repository_test

import (
	"simushop/internal/core"
	"simushop/internal/entity"
	"testing"
)

func TestCreatingProduct(t *testing.T) {
	t.Log("Creating Product")

	err := createProduct(entity.Product{
		ProductName:             "Produto de Teste teste",
		ProductValue:            123,
		ProductDiscountPercent:  2,
		ProductDescription:      "Descrição de teste",
		ProductKgWeight:         123,
		ProductCentimeterWidth:  123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 35,
		ProductSellerId:         1,
	})

	equal(t, nil, err, "Error on saving user on Database.")
}

func TestTryCreatingProductsWithSameName(t *testing.T) {
	t.Log("testing creating products with same names.")

	product := entity.Product{
		ProductName:             "Produto de Teste",
		ProductValue:            123,
		ProductDiscountPercent:  2,
		ProductDescription:      "Descrição de teste",
		ProductKgWeight:         123,
		ProductCentimeterWidth:  123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 35,
		ProductSellerId:         1,
	}

	/* creating product*/
	createProduct(product)

	/* creating product with same name*/
	err := createProduct(product)

	notEqual(t, nil, err, "Saving user with same names into the database.")
}

func TestUpdateProduct(t *testing.T) {
	t.Log("testing update product")
	product := entity.Product{
		ProductName:             "Produto de Teste",
		ProductValue:            123,
		ProductDiscountPercent:  2,
		ProductDescription:      "Descrição de teste",
		ProductKgWeight:         123,
		ProductCentimeterWidth:  123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 35,
		ProductSellerId:         1,
	}
	productNameToUpdate := product.ProductName + "novo"

	createProduct(product)

	err := mockRepository.UpdateProduct(entity.Product{
		ProductName:  productNameToUpdate,
		ProductValue: 888,
	}, "product_name = ? ", productNameToUpdate)

	equal(t, nil, err, "error on update Product.")
}

func TestListProduct(t *testing.T) {
	t.Log("testing list Products")

	for _, i := range []string{"1", "2", "3"} {
		createProduct(entity.Product{
			ProductName:             "Produto de Teste listagem " + i,
			ProductValue:            123,
			ProductDiscountPercent:  2,
			ProductDescription:      "Descrição de teste",
			ProductKgWeight:         123,
			ProductCentimeterWidth:  123,
			ProductCentimeterHeight: 123,
			ProductCentimeterLength: 35,
			ProductSellerId:         1,
		})
	}

	products := mockRepository.PaginateTopics(core.NewPaginate(core.Request{
		GetQueryParam: func(key string) string {
			if key == "page" {
				return "1"
			} else {
				return "3"
			}
		},
	}), "")

	equal(t, true, len(products) == 3, "error on list products.")
}

func createProduct(e entity.Product) error {
	return mockRepository.CreateProduct(e)
}

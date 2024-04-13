package repository

import (
	"simushop/internal/entity"
)

func (r repository) CreateProduct(product entity.Product) error {
	var (
		productFound entity.Product
		result       error = newError("product with name '" + product.ProductName + "' exists.")
	)

	r.queryFirst(&productFound, "product_name = ?", product.ProductName)

	if !productFound.Exist() {
		result = r.db.Create(&product).Error
	}

	return result
}

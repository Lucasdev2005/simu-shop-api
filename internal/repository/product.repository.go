package repository

import (
	"simushop/internal/entity"
)

func (r repository) CreateProduct(product entity.Product) error {
	var (
		result error = newError("product with name '" + product.ProductName + "' exists.")
	)

	if !productNameExists(product.ProductName, r) {
		result = r.db.Create(&product).Error
	}

	return result
}

func (r repository) UpdateProduct(product entity.Product, where string, args ...interface{}) error {
	var (
		result      error = nil
		applyUpdate       = true
	)

	if len(product.ProductName) > 0 && productNameExists(product.ProductName, r) {
		applyUpdate = false
		result = newError("product with name '" + product.ProductName + "' exists.")
	}

	if applyUpdate {
		result = r.update(&product, where, args...).Error
	}

	return result
}

func productNameExists(productName string, r repository) bool {
	var productFound entity.Product

	r.queryFirst(&productFound, "product_name = ?", productName)

	return productFound.Exist()
}
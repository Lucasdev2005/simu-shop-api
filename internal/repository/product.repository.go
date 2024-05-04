package repository

import (
	"simushop/internal/core"
	"simushop/internal/entity"
)

func (r repository) CreateProduct(product entity.Product) (entity.Product, error) {
	if !productNameExists(product.ProductName, r) {
		err := r.db.Create(&product).Error
		if err != nil {
			return entity.Product{}, err
		}
	}

	return product, nil
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

func (r repository) PaginateTopics(paginate core.Paginate, where string, args ...interface{}) []entity.Product {
	var products []entity.Product

	if len(where) > 0 {
		r.paginate(paginate).Preload("Topics").Where(where, args...).Find(&products)
	} else {
		r.paginate(paginate).Find(&products)
	}

	return products
}

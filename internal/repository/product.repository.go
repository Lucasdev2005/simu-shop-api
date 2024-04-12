package repository

import "simushop/internal/entity"

func (r repository) CreateProduct(product entity.Product) error {
	return r.db.Create(&product).Error
}

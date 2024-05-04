package repository

import (
	"log"
	"simushop/internal/entity"

	"gorm.io/gorm"
)

func (r repository) AddItemOncart(shoppingCart entity.ShoppingCart) {
	r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Save(&shoppingCart).Error
	})

	log.Println("[AddItemOnCart] shoppingCart: ", shoppingCart)
}

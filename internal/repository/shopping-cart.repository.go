package repository

import (
	"fmt"
	"simushop/internal/entity"

	"gorm.io/gorm"
)

func (r repository) AddItemOnCart(shoppingCart entity.ShoppingCart) (entity.ShoppingCart, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var shoppingCartItemFound entity.ShoppingCart
		r.db.Where(
			"shopping_cart_item_id = ? AND shopping_cart_user_id = ?",
			shoppingCart.ShoppingCartItemId,
			shoppingCart.ShoppingCartUserId,
		).Find(&shoppingCartItemFound)

		if shoppingCartItemFound.ShoppingCartId != 0 {
			return fmt.Errorf("product exists on cart")
		}

		errOnSave := tx.Create(&shoppingCart).Error

		return errOnSave
	})

	if err != nil {
		return entity.ShoppingCart{}, err
	} else {
		return shoppingCart, nil
	}
}

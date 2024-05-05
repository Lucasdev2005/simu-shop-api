package dto

import "simushop/internal/entity"

type AddItemCartDTO struct {
	ShoppingCartItemQuantity int `json:"shoppingCartItemQuantity,omitempty" validate:"min=1,required"`
	ShoppingCartItemId       int `json:"shoppingCartItemId,omitempty" validate:"min=1,required"`
	ShoppingCartUserId       int `json:"shoppingCartUserId,omitempty" validate:"min=1,required"`
}

func (a AddItemCartDTO) ToShoppingCart() entity.ShoppingCart {
	return entity.ShoppingCart{
		ShoppingCartItemQuantity: a.ShoppingCartItemQuantity,
		ShoppingCartItemId:       a.ShoppingCartItemId,
		ShoppingCartUserId:       a.ShoppingCartUserId,
	}
}

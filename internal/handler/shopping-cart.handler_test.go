package handler_test

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"testing"
)

func TestAddItemOnShoppingCart(t *testing.T) {
	t.Log("testing add item on shopping cart")

	res := AddItemOnCart(dto.AddItemCartDTO{
		ShoppingCartItemQuantity: 6,
		ShoppingCartItemId:       2,
		ShoppingCartUserId:       1,
	})

	Created(res, t)
}

func AddItemOnCart(data dto.AddItemCartDTO) core.Response {
	return handlerInstance.AddItemOnCart(core.Request{
		Body: func(obj any) error {
			addItemDTO, _ := obj.(*dto.AddItemCartDTO)

			addItemDTO.ShoppingCartItemQuantity = data.ShoppingCartItemQuantity
			addItemDTO.ShoppingCartItemId = data.ShoppingCartItemId
			addItemDTO.ShoppingCartUserId = data.ShoppingCartUserId

			return nil
		},
	})
}

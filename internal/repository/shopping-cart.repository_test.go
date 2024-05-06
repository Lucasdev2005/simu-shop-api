package repository_test

import (
	"simushop/internal/entity"
	"simushop/internal/enum"
	"testing"
)

func TestAddItemOnCart(t *testing.T) {
	t.Log("Add item on shopping Cart")

	user, _ := CreateUser(entity.User{
		UserPassword: "gerinus@123",
		Username:     "Lucas Teste creating Item",
		UserBalance:  300,
		UserType:     enum.UserTypeEnum.BUYER,
	})

	topic, _ := CreateTopic(entity.Topic{
		TopicName: "Topico de teste para adicionar item no carrinho",
	})

	product, _ := CreateProduct(entity.Product{
		ProductName:             "Produto de Teste",
		ProductValue:            123,
		ProductDiscountPercent:  2,
		ProductDescription:      "Descrição de teste",
		ProductKgWeight:         123,
		ProductCentimeterWidth:  123,
		ProductCentimeterHeight: 123,
		ProductCentimeterLength: 35,
		ProductSellerId:         1,
		Topics:                  []entity.Topic{{TopicId: topic.TopicId}},
	})

	mockRepository.AddItemOnCart(entity.ShoppingCart{
		ShoppingCartItemQuantity: 1,
		ShoppingCartItemId:       product.ProductId,
		ShoppingCartUserId:       user.UserId,
	})
}

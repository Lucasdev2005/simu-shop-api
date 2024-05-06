package repository_test

import (
	"math/rand"
	"simushop/internal/entity"
	"simushop/internal/enum"
	"testing"
	"time"
)

func TestAddItemOnCart(t *testing.T) {
	t.Log("Add item on shopping Cart")

	_, err := createAndAddItemOnCart()

	equal(t, nil, err, "Error when Adding item to Cart")
}

func TestRemoveItemOnCart(t *testing.T) {
	t.Log("removing item from Cart")

	item, _ := createAndAddItemOnCart()

	err := mockRepository.RemoveItemOnCart(item.ShoppingCartItemId, item.ShoppingCartUserId)

	equal(t, nil, err, "Error on Remove item from Cart")
}

func createAndAddItemOnCart() (entity.ShoppingCart, error) {
	user, _ := CreateUser(entity.User{
		UserPassword: "gerinus@123",
		Username:     randStringRunes(9),
		UserBalance:  300,
		UserType:     enum.UserTypeEnum.BUYER,
	})

	topic, _ := CreateTopic(entity.Topic{
		TopicName: randStringRunes(9),
	})

	product, _ := CreateProduct(entity.Product{
		ProductName:             randStringRunes(9),
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

	return mockRepository.AddItemOnCart(entity.ShoppingCart{
		ShoppingCartItemQuantity: 1,
		ShoppingCartItemId:       product.ProductId,
		ShoppingCartUserId:       user.UserId,
	})
}

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

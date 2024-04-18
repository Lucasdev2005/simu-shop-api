package handler_test

import (
	"fmt"
	"simushop/internal/entity"
	"strconv"
)

type mockRepositoryImpl struct{}

func (m mockRepositoryImpl) CreateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (m mockRepositoryImpl) UpdateUser(user entity.User, where string, args ...interface{}) error {
	users := map[string]entity.User{}

	for _, element := range []int{1, 4, 5} {
		users[strconv.Itoa(element)] = entity.User{
			UserId:       1,
			UserPassword: "teste@!23",
			Username:     "Lucas dos teste" + strconv.Itoa(element),
			UserBalance:  1234,
			UserType:     "seller",
		}
	}

	if _, ok := users[args[0].(string)]; !ok {
		return fmt.Errorf(args[0].(string) + " not found")
	}

	return nil
}

func (m mockRepositoryImpl) CreateProduct(product entity.Product) error {

	if productTable()[product.ProductName].Exist() {
		return fmt.Errorf("Product " + product.ProductName + "Exists.")
	}

	return nil
}

func (m mockRepositoryImpl) UpdateProduct(product entity.Product, where string, args ...interface{}) error {

	if len(product.ProductName) > 0 && productTable()[product.ProductName].Exist() {
		return fmt.Errorf("Product " + product.ProductName + " Exists.")
	}

	return nil
}

func (m mockRepositoryImpl) CreateTopic(topic entity.Topic) error {
	var (
		table        = topicTable()
		result error = nil
	)

	if _, ok := table[topic.TopicName]; ok {
		result = fmt.Errorf("topic exists")
	}

	return result
}

func (m mockRepositoryImpl) UpdateTopic(topic entity.Topic, where string, args ...interface{}) error {
	return nil
}

func (m mockRepositoryImpl) ListTopics(where string, args ...interface{}) []entity.Topic {

	table := topicTable()

	var topics []entity.Topic

	for _, topic := range table {
		topics = append(topics, topic)
	}

	return topics
}

func productTable() map[string]entity.Product {
	products := map[string]entity.Product{}

	for _, element := range []string{"Lucas Moreira", "produto de teste", "produto de teste 2"} {
		products[element] = entity.Product{
			ProductId:               3,
			ProductName:             element,
			ProductValue:            123,
			ProductDiscountPercent:  1,
			ProductDescription:      "descrição de teste",
			ProductKgWeight:         23,
			ProductCentimeterWidth:  120,
			ProductCentimeterHeight: 11,
			ProductCentimeterLength: 11,
		}
	}

	return products
}

func topicTable() map[string]entity.Topic {
	var topics = map[string]entity.Topic{}

	for _, i := range []int{1, 2, 3, 4, 5} {
		topics["topic "+strconv.Itoa(i)] = entity.Topic{
			TopicId:   i,
			TopicName: "topic " + strconv.Itoa(i),
		}
	}

	return topics
}

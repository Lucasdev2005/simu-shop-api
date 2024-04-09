package handler_test

import (
	"fmt"
	"net/http"
	"simushop/internal/core"
	"simushop/internal/entity"
	"simushop/internal/handler"
	"strconv"
	"testing"
)

type mockRepositoryImpl struct{}

func (m *mockRepositoryImpl) CreateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (m *mockRepositoryImpl) UpdateUser(user entity.User, where string, args ...interface{}) error {
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

var (
	mockRepository  = &mockRepositoryImpl{}
	handlerInstance = handler.NewHandler(mockRepository)
)

func BadRequest(result core.Response, t *testing.T) {
	if result.Code != http.StatusBadRequest {
		t.Error(result.Data)
		t.Errorf("response Code Should be a 400 bad Request but got " + strconv.Itoa(result.Code))
	}
}

func Created(result core.Response, t *testing.T) {
	if result.Code != http.StatusCreated {
		t.Error(result.Data)
		t.Errorf("response Code Should be a 201 Created but got " + strconv.Itoa(result.Code))
	}
}

func Ok(result core.Response, t *testing.T) {
	if result.Code != http.StatusOK {
		t.Error(result.Data)
		t.Errorf("response Code Should be a 200 Created but got " + strconv.Itoa(result.Code))
	}
}

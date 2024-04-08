package repository_test

import (
	"simushop/internal/entity"
	"simushop/internal/enum"
	"testing"
)

var user entity.User = entity.User{
	UserPassword: "gerinus@123",
	Username:     "lucas moreira nunes",
	UserBalance:  300,
	UserType:     enum.UserTypeEnum.BUYER,
}

func TestCreateUser(t *testing.T) {
	err := mockRepository.CreateUser(user)
	t.Error(err)
}

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
	t.Log("testing creating a user.")
	err := createUser()

	if err != nil {
		t.Error(err)
	}
}

func TestCreatingUsernameWithExistingUsername(t *testing.T) {
	t.Log("testing creating user with username existing on database.")
	createUser()
	err := createUser()

	if err == nil {
		t.Error("saving user with a existing username.")
	}
}

func createUser() error {
	return mockRepository.CreateUser(user)
}

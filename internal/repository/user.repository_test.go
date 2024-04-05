package repository_test

import (
	"simushop/internal/entity"
	"simushop/internal/enum"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Log("testing create a user.")

	sqlMock.ExpectQuery(`^SELECT (.+) FROM "user" WHERE username =`).
		WithArgs("lucas moreira nunes", 1)

	sqlMock.ExpectBegin()

	mockRepository.CreateUser(entity.User{
		UserPassword: "gerinus@123",
		Username:     "lucas moreira nunes",
		UserBalance:  300,
		UserType:     enum.UserTypeEnum.BUYER,
	})

	success(t)
}

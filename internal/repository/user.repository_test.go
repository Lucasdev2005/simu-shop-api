package repository_test

import (
	"simushop/internal/entity"
	"simushop/internal/enum"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Log("testing creating a user.")

	_, err := createUser(entity.User{
		UserPassword: "gerinus@123",
		Username:     "lucas moreira nunes",
		UserBalance:  300,
		UserType:     enum.UserTypeEnum.BUYER,
	})

	equal(t, nil, err, "error on saving user.")
}

func TestCreatingUsernameWithExistingUsername(t *testing.T) {
	t.Log("testing creating user with username existing on database.")
	user := entity.User{
		UserPassword: "gerinus@123",
		Username:     "Lucas Teste existing name",
		UserBalance:  300,
		UserType:     enum.UserTypeEnum.BUYER,
	}

	// creating first user with 'lucas moreira nunes' name.
	createUser(user)

	_, err := createUser(user)
	notEqual(t, nil, err, "saving user with a existing username.")
}

func TestUpdatingUser(t *testing.T) {
	t.Log("testing update user.")

	userCreated, err := createUser(entity.User{
		UserPassword: "gerinus@123",
		Username:     "Lucas Teste update name",
		UserBalance:  300,
		UserType:     enum.UserTypeEnum.BUYER,
	})
	equal(t, nil, err, "error on saving user.")

	t.Error("[TestUpdatingUser] userCreated.UserId", userCreated.UserId)
	errOnUpdate := mockRepository.UpdateUser(entity.User{
		Username: "novo lucas moreira nunes",
	}, "user_id = ?", userCreated.UserId)

	equal(t, nil, errOnUpdate, "error on updating user: "+strconv.Itoa(userCreated.UserId))
}

func createUser(user entity.User) (entity.User, error) {
	return mockRepository.CreateUser(user)
}

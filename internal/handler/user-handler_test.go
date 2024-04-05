package handler_test

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"simushop/internal/enum"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Log("Testing creating a user.")

	result := createUser(dto.CreateUserDTO{
		Username:     "exemplo_usuario",
		UserBalance:  100.00,
		UserPassword: "senha_segura",
		UserType:     enum.UserTypeEnum.SELLER,
	}, t)

	Created(result, t)
}

func TestRequiredFieldsFromCreateUser(t *testing.T) {
	t.Log("Testing required fields from create a User.")

	result := createUser(dto.CreateUserDTO{
		Username:     "",
		UserPassword: "",
		UserBalance:  1,
		UserType:     "",
	}, t)

	BadRequest(result, t)
}

func TestNegativeBalanceFromCreateUer(t *testing.T) {
	t.Log("Testing negative balance on create a User")

	result := createUser(dto.CreateUserDTO{
		Username:     "exemplo_usuario",
		UserBalance:  -100.4,
		UserPassword: "senha_segura",
		UserType:     enum.UserTypeEnum.SELLER,
	}, t)

	BadRequest(result, t)
}

func createUser(createUserDTO dto.CreateUserDTO, t *testing.T) core.Response {
	return handlerInstance.CreateUser(core.Request{
		Body: func(obj any) error {
			createUserDTOFromTest, ok := obj.(*dto.CreateUserDTO)

			if !ok {
				t.Error("Invalid 'CreateUserDTO'")
			}

			createUserDTOFromTest.Username = createUserDTO.Username
			createUserDTOFromTest.UserBalance = createUserDTO.UserBalance
			createUserDTOFromTest.UserPassword = createUserDTO.UserPassword
			createUserDTOFromTest.UserType = createUserDTO.UserType

			return nil
		},
	})
}

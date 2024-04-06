package repository_test

import (
	"regexp"
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
	t.Log("testing create a User.")

	sqlMock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "user" WHERE username = $1 ORDER BY "user"."user_id" LIMIT $2`),
	).WithArgs(user.Username, 1)

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(`INSERT INTO "user" (.+) VALUES (.+) RETURNING "user_id"`).WithArgs(user.UserPassword, user.Username, user.UserBalance, user.UserType)

	mockRepository.CreateUser(user)

	success(t)
}

func TestUpdateUser(t *testing.T) {
	t.Log("testing update a User.")

	user.UserId = 1
	user.Username = "teste update"
	mockRepository.UpdateUser(user, "user_id = $1", user.UserId)

	success(t)
}

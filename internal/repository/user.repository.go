package repository

import (
	"fmt"
	"log"
	"simushop/internal/entity"
)

func (r repository) CreateUser(user entity.User) (entity.User, error) {
	var (
		userFinded entity.User

		result error = fmt.Errorf("username " + user.Username + " exists.")
	)

	r.queryFirst(&userFinded, "username = $1", user.Username)

	if !userFinded.Exist() {
		result = r.db.Create(&user).Error
	}

	return user, result
}

func (r repository) UpdateUser(user entity.User, where string, args ...interface{}) error {
	var result error = nil

	fmt.Println("[UpdateUser] where, args: ", args)
	fmt.Println("[UpdateUser] where, args: ", where)
	tx := r.db.Model(&user).Select("*").Where(where, args...).Updates(map[string]interface{}{
		"username": user.Username,
	})

	log.Println("[UpdateUser] user: ", user)

	if tx.Error != nil {
		result = tx.Error
	}

	log.Println("[UpdateUser] result: ", result)
	return result
}

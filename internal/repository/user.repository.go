package repository

import (
	"fmt"
	"simushop/internal/entity"
)

func (r repository) CreateUser(user entity.User) (entity.User, error) {
	var (
		userFound entity.User

		result error = fmt.Errorf("username " + user.Username + " exists.")
	)

	r.queryFirst(&userFound, "username = $1", user.Username)

	if !userFound.Exist() {
		result = r.db.Create(&user).Error
	}

	return user, result
}

func (r repository) UpdateUser(user entity.User, where string, args ...interface{}) error {
	var result error = nil

	tx := r.update(&user, where, args...)

	if tx.Error != nil {
		result = tx.Error
	}

	return result
}

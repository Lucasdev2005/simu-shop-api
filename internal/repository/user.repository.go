package repository

import (
	"simushop/internal/entity"
)

func (r repository) CreateUser(user entity.User) (entity.User, error) {
	var (
		userFound entity.User

		result error = newError("username " + user.Username + " exists.")
	)

	r.queryFirst(&userFound, "username = $1", user.Username)

	if !userFound.Exist() {
		user.SetPassword()
		result = r.db.Create(&user).Error
	}

	return user, result
}

func (r repository) UpdateUser(user entity.User, where string, args ...interface{}) error {
	var result error = nil

	if len(user.UserPassword) > 0 {
		user.SetPassword()
	}

	tx := r.update(&user, where, args...)

	if tx.Error != nil {
		result = tx.Error
	}

	return result
}

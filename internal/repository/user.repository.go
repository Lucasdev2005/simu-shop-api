package repository

import (
	"fmt"
	"simushop/internal/entity"
)

func (r repository) CreateUser(user entity.User) error {
	var (
		userFinded entity.User

		result error = fmt.Errorf("username " + user.Username + " exists.")
	)

	r.queryFirst(&userFinded, "username = $1", user.Username)

	if !userFinded.Exist() {
		result = r.db.Create(&user).Scan(&user).Error
	}

	return result
}

func (r repository) UpdateUser(user entity.User, where string, args ...interface{}) {
	r.db.Model(&user).Where(where, args).Updates(r.mapFields(user))
}

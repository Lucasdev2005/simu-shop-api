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

	r.db.Where("username = $1", user.Username).First(&userFinded)

	if !userFinded.Exist() {
		result = r.db.Create(&user).Scan(&user).Error
	}

	return result
}

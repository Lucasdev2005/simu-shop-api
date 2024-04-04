package repository

import (
	"fmt"
	"simushop/internal/dto"
	"simushop/internal/entity"
)

func (r repository) CreateUser(createUserDTO dto.CreateUserDTO) error {
	var (
		userFinded entity.User
		user       = entity.User{
			UserPassword: createUserDTO.UserPassword,
			Username:     createUserDTO.Username,
			UserBalance:  createUserDTO.UserBalance,
			UserType:     createUserDTO.UserType,
		}

		result error = fmt.Errorf("username " + user.Username + " exists.")
	)

	r.db.Where("username = $1", user.Username).First(&userFinded)

	if !userFinded.Exist() {
		result = r.db.Save(&user).Scan(&user).Error
	}

	return result
}

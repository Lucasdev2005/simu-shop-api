package repository

import "simushop/internal/entity"

func (r repository) CreateUser(user entity.User) {
	r.db.Save(user)
}

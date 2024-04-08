package repository_test

import (
	"simushop/internal/database"
	"simushop/internal/entity"
	"simushop/internal/repository"

	"gorm.io/driver/sqlite"
)

var db = database.NewDatabase(
	sqlite.Open("gorm.db"),
	entity.User{},
)

var mockRepository = repository.NewRepository(db.Db)

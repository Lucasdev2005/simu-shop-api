package repository_test

import (
	"simushop/internal/database"
	"simushop/internal/entity"
	"simushop/internal/repository"
	"testing"

	"gorm.io/driver/sqlite"
)

var db = database.NewDatabase(
	sqlite.Open("gorm.db"),
	entity.User{},
)

var mockRepository = repository.NewRepository(db.Db)

/* test if two values are equal. */
func equal(t *testing.T, expected interface{}, actual interface{}, msg string) {
	if expected != actual {
		t.Error(msg)
	}
}

/* test if two values are not equal. */
func notEqual(t *testing.T, expected interface{}, actual interface{}, msg string) {
	if expected == actual {
		t.Error(msg)
	}
}

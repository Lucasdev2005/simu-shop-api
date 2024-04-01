package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	Db *gorm.DB
}

func NewDatabase(port string, password string, user string, name string, host string, entities ...interface{}) database {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		name,
		port,
	)), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	if len(entities) > 0 {
		db.AutoMigrate(entities...)
	}

	return database{
		db,
	}
}

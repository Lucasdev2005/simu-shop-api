package database

import (
	"log"

	"gorm.io/gorm"
)

type database struct {
	Db *gorm.DB
}

func NewDatabase(dialector gorm.Dialector, entities ...interface{}) database {
	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})

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

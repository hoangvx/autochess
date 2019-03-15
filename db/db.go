package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
		"autochess/entity"
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=app_user dbname=autochess password=admin sslmode=disable")
    if err != nil {
        panic(err)
    }

		autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}

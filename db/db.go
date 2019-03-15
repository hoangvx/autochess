package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "autochess/entity"
    "os"
)

var (
    db  *gorm.DB
    err error
)

// const LOCAL_URL = "postgres://app_user:admin@localhost:5432/autochess?sslmode=disable"

// Init is initialize db from main function
func Init() {
    db_url := os.Getenv("DATABASE_URL")
    // if db_url == "" {
    //     db_url = LOCAL_URL
    // }
    db, err = gorm.Open("postgres", db_url)
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

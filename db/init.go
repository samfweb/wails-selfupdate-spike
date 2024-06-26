package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"wails-selfupdate-spike/models"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if res := db.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
		panic("failed to enable foreign_keys: " + res.Error.Error())
	}

	db.AutoMigrate(&models.Connection{})
	db.AutoMigrate(&models.Subscription{})

	return db
}

package util

import (
	"user_ms/backend/core/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {

	if DB == nil {
		db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

		if err != nil {
			panic("failed to connect to database")
		}

		// Auto Migrate
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			panic("failed to migrate database")
		}

		DB = db
	}

	return DB
}

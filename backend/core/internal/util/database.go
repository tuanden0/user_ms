package util

import (
	"fmt"
	"sync"
	"user_ms/backend/core/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	DBErr      error
	dbConnOnce sync.Once
)

func ConnectDatabase() (*gorm.DB, error) {

	dbConnOnce.Do(func() {
		db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

		if err != nil {
			DBErr = fmt.Errorf("failed to connect to database")
		}

		// Auto Migrate
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			DBErr = fmt.Errorf("failed to migrate database")
		}

		DB = db
	})

	if DBErr != nil {
		return nil, DBErr
	}

	return DB, nil
}

package db

import (
	"log"

	"github.com/Anirudh-RedLion/student_pm_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection using the provided dbPath.
func Init(dbPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate User model
	DB.AutoMigrate(&models.User{})
}

package db

import (
	"log"
	"student_pm_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("studentpm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate User model
	DB.AutoMigrate(&models.User{})
}

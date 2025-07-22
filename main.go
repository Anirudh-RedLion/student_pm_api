package main

import (
	"log"
	"os"

	"github.com/Anirudh-RedLion/student_pm_api/db"
	"github.com/Anirudh-RedLion/student_pm_api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present (for local dev)
	_ = godotenv.Load()

	// Get configuration from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "studentpm.db" // default DB file
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	// Pass config to db and utils packages as needed
	db.Init(dbPath)

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":" + port)
}

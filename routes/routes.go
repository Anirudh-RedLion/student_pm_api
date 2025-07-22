package routes

import (
	"student_pm_api/handlers"
	"student_pm_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/users/me", middleware.JWTAuth(), handlers.GetMe)
	}
}

package routes

import (
  "github.com/gin-gonic/gin"
  "rewarding-harvest-backend/controllers"
)

// SetupRoutes - Register Routes
func SetupRoutes(router *gin.Engine) {
  api := router.Group("/api")
  {
    api.GET("/health", controllers.HealthCheck)
    api.POST("/auth/register", controllers.RegisterUser)
  }
}

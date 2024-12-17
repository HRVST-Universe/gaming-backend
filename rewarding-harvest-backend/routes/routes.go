package routes

import (
  "github.com/gin-gonic/gin"
  "rewarding-harvest-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
  api := router.Group("/api")
  {
    // User Routes
    api.POST("/users/register", controllers.RegisterUser)
    api.POST("/auth/login", controllers.LoginUser)

    // Player Actions Routes
    api.POST("/player-actions", controllers.LogPlayerAction)

    // Health Check Route
    api.GET("/health", func(c *gin.Context) {
      c.JSON(200, gin.H{"status": "Server is running"})
    })
  }
}

package routes

import (
  "github.com/gin-gonic/gin"
  "rewarding-harvest-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
  api := router.Group("/api")
  {
    // User Management Endpoints
    api.POST("/users/register", controllers.RegisterUser)
    api.POST("/users/login", controllers.LoginUser)
    api.GET("/users", controllers.GetAllUsers)
    api.GET("/users/email/:email", controllers.GetUserByEmail)
  api.GET("/users/wallet/:walletAddress", controllers.GetUserByWalletAddress)
    api.PUT("/users/email/:email", controllers.UpdateUserByEmail)

    // Player Actions Endpoints
    api.POST("/player-actions", controllers.LogPlayerAction)
    api.GET("/player-actions/id/:playerId", controllers.GetPlayerActionsByID)
    api.GET("/player-actions/email/:playerEmail", controllers.GetPlayerActionsByEmail)
    api.GET("/player-actions/wallet/:playerWallet", controllers.GetPlayerActionsByWallet)

    // Health Check Endpoint
    api.GET("/health", func(c *gin.Context) {
      c.JSON(200, gin.H{"status": "OK"})
    })
  }
}

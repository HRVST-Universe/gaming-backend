package main

import (
  "log"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/routes"
)

func main() {
  // Load environment variables
  if err := godotenv.Load(".env"); err != nil {
    log.Println("⚠️  Warning: .env file not found.")
  }

  // Initialize the database
  config.ConnectDatabase()

  // Initialize the router
  router := gin.Default()

  // Setup routes
  routes.SetupRoutes(router)

  // Start the server
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  log.Printf("🚀 Server running on port %s", port)
  if err := router.Run(":" + port); err != nil {
    log.Fatalf("❌ Server failed: %v", err)
  }
}

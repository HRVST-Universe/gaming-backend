package main

import (
  "log"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/controllers"
)

func main() {
  // Load environment variables
  err := godotenv.Load(".env")
  if err != nil {
    log.Println("Warning: .env file not found. Using system environment variables.")
  }

  // Set Gin mode
  ginMode := os.Getenv("GIN_MODE")
  if ginMode == "" {
    ginMode = gin.ReleaseMode
  }
  gin.SetMode(ginMode)

  // Connect to the database
  config.ConnectDatabase()

  // Initialize Gin server
  r := gin.Default()

  // Set trusted proxies
  err = r.SetTrustedProxies(nil)
  if err != nil {
    log.Fatalf("Failed to set trusted proxies: %v", err)
  }

  // Register Root Route
  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Welcome to Rewarding Harvest API",
    })
  })

  // Register API Routes
  api := r.Group("/api")
  {
    api.GET("/health", controllers.HealthCheck)         // GET /api/health
    api.POST("/auth/register", controllers.RegisterUser) // POST /api/auth/register
  }

  // Start the server
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }
  log.Printf("Starting server on port %s", port)
  if err := r.Run(":" + port); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}

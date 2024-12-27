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
  log.Println("🚀 Starting Rewarding Harvest Backend...")

  // Load environment variables
  if err := godotenv.Load(".env"); err != nil {
    log.Println("⚠️  Warning: .env file not found. Using system environment variables.")
  }

  // Connect to the database
  log.Println("🔌 Connecting to the database...")
  config.ConnectDatabase()
  log.Println("✅ Database connected successfully")

  // Initialize the router
  router := gin.Default()

  // Configure CORS
  router.Use(func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
    }
    c.Next()
  })

  // Register routes
  log.Println("🛠 Registering routes...")
  routes.SetupRoutes(router)

  // Root Health Check for Deployment
  router.GET("/", func(c *gin.Context) {
    log.Println("✅ Root health check passed.")
    c.JSON(200, gin.H{
      "status":  "OK",
      "message": "Rewarding Harvest Backend is running",
    })
  })

  // Global 404 Handler
  router.NoRoute(func(c *gin.Context) {
    log.Printf("❌ 404 - Path not found: %s %s", c.Request.Method, c.Request.URL.Path)
    c.JSON(404, gin.H{"error": "Resource not found"})
  })

  // Start Server
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }
  log.Printf("✅ Server running on port %s", port)

  if err := router.Run("0.0.0.0:" + port); err != nil {
    log.Fatalf("❌ Server startup failed: %v", err)
  }
}

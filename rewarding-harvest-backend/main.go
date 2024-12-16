package main

import (
  "log"
  "os"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"

  // Internal packages
  "github.com/HRVST-Universe/gaming-backend/rewarding-harvest-backend/config"
  "github.com/HRVST-Universe/gaming-backend/rewarding-harvest-backend/routes"
  "github.com/HRVST-Universe/gaming-backend/rewarding-harvest-backend/schemas"
  "github.com/HRVST-Universe/gaming-backend/rewarding-harvest-backend/utils"
)

func main() {
  // Load environment variables
  if err := godotenv.Load(".env"); err != nil {
    log.Println("⚠️  Warning: .env file not found. Using system environment variables.")
  }

  // Connect to the database
  config.ConnectDatabase()

  // Fetch PostgreSQL schema and generate models
  schema := schemas.FetchSchema(config.DB)
  utils.GenerateModels(schema)
  log.Println("✅ Models generated successfully")

  // Initialize Gin router
  r := gin.Default()

  // Health check route
  r.GET("/api/health", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status":  "Server is running",
      "message": "Welcome to Rewarding Harvest API",
    })
  })

  // API Discovery Route: List all available routes
  r.GET("/api/routes", func(c *gin.Context) {
    routesInfo := []gin.RouteInfo{}
    for _, route := range r.Routes() {
      routesInfo = append(routesInfo, route)
    }
    c.JSON(200, gin.H{
      "status": "success",
      "routes": routesInfo,
    })
  })

  // Auto-generate CRUD routes
  var models []string
  for tableName := range utils.GroupSchemaByTable(schema) {
    models = append(models, tableName)
  }
  routes.AutoGenerateRoutes(r, config.DB, models)
  log.Println("✅ CRUD routes registered successfully")

  // Start the server
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  log.Printf("🚀 Server running on port %s", port)
  if err := r.Run(":" + port); err != nil {
    log.Fatalf("❌ Failed to start the server: %v", err)
  }
}

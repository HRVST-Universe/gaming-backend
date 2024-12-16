package routes

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

// AutoGenerateRoutes registers CRUD routes for all models
func AutoGenerateRoutes(r *gin.Engine, DB *gorm.DB, models []string) {
  for _, model := range models {
    r.GET(fmt.Sprintf("/api/%s", model), func(c *gin.Context) {
      var results []interface{}
      if err := DB.Table(model).Find(&results).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve records"})
        return
      }
      c.JSON(http.StatusOK, results)
    })

    r.POST(fmt.Sprintf("/api/%s", model), func(c *gin.Context) {
      var payload map[string]interface{}
      if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
      }

      if err := DB.Table(model).Create(&payload).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
        return
      }

      c.JSON(http.StatusCreated, gin.H{"status": "success"})
    })
  }
}

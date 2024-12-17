package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/models"
)

// Log Player Action
func LogPlayerAction(c *gin.Context) {
  var action models.PlayerAction

  if err := c.ShouldBindJSON(&action); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  if action.ActionType == "" || action.PlayerID == "" || action.PlayerEmail == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
    return
  }

  if err := config.DB.Create(&action).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Action logging failed"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "Action logged successfully", "action": action})
}

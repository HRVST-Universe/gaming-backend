package controllers

import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/models"
)

// Log Player Action (POST /api/player-actions)
func LogPlayerAction(c *gin.Context) {
  var payload models.PlayerAction

  // Validate request payload
  if err := c.ShouldBindJSON(&payload); err != nil {
    log.Printf("❌ Invalid input: %v", err)
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  // Save action to the database
  if err := config.DB.Create(&payload).Error; err != nil {
    log.Printf("❌ Failed to log action: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log action"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"success": true,"message": "Action logged successfully", "action": payload})
}

// Get Player Actions by Player ID (GET /api/player-actions/id/:playerId)
func GetPlayerActionsByID(c *gin.Context) {
  playerID := c.Param("playerId")

  var actions []models.PlayerAction
  if err := config.DB.Where("player_id = ?", playerID).Find(&actions).Error; err != nil {
    log.Printf("❌ Actions for player ID %s not found: %v", playerID, err)
    c.JSON(http.StatusNotFound, gin.H{"error": "Actions not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"actions": actions})
}

// Get Player Actions by Email (GET /api/player-actions/email/:playerEmail)
func GetPlayerActionsByEmail(c *gin.Context) {
  playerEmail := c.Param("playerEmail")

  var actions []models.PlayerAction
  if err := config.DB.Where("player_email = ?", playerEmail).Find(&actions).Error; err != nil {
    log.Printf("❌ Actions for email %s not found: %v", playerEmail, err)
    c.JSON(http.StatusNotFound, gin.H{"error": "Actions not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"actions": actions})
}

// Get Player Actions by Wallet Address (GET /api/player-actions/wallet/:playerWallet)
func GetPlayerActionsByWallet(c *gin.Context) {
  playerWallet := c.Param("playerWallet")

  var actions []models.PlayerAction
  if err := config.DB.Where("player_wallet = ?", playerWallet).Find(&actions).Error; err != nil {
    log.Printf("❌ Actions for wallet %s not found: %v", playerWallet, err)
    c.JSON(http.StatusNotFound, gin.H{"error": "Actions not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"actions": actions})
}

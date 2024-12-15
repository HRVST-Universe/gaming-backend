package controllers

import (
  "net/http"
  "os"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v5"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/models"
)

// Health Check Endpoint
func HealthCheck(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": "Server running"})
}

// Register User Endpoint
func RegisterUser(c *gin.Context) {
  var user models.User

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  if err := config.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    return
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":    user.ID,
    "email": user.Email,
    "exp":   time.Now().Add(time.Hour * 1).Unix(),
  })

  tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status":  "success",
    "message": "User registered successfully",
    "token":   tokenString,
  })
}

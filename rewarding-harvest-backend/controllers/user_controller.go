package controllers

import (
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/models"
  "github.com/golang-jwt/jwt/v5"
)

// Register User
func RegisterUser(c *gin.Context) {
  var payload struct {
    GameShiftID string `json:"gameshiftId"`
    Email       string `json:"email"`
    Username    string `json:"username"`
  }

  if err := c.ShouldBindJSON(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  user := models.User{
    GameShiftID:   payload.GameShiftID,
    Email:         payload.Email,
    Username:      payload.Username,
    WalletType:    "unknown",
    WalletAddress: "unknown",
  }

  if err := config.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login User
func LoginUser(c *gin.Context) {
  var payload struct {
    Email string `json:"email"`
  }

  if err := c.ShouldBindJSON(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  var user models.User
  if err := config.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":    user.ID,
    "email": user.Email,
    "exp":   time.Now().Add(time.Hour * 24).Unix(),
  })

  tokenString, err := token.SignedString([]byte("SECRET_KEY"))
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Login successful",
    "user":    user,
    "token":   tokenString,
  })
}

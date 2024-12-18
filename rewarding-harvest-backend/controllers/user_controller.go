package controllers

import (
  "log"
  "net/http"
  "os"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v5"
  "rewarding-harvest-backend/config"
  "rewarding-harvest-backend/models"
)

// Register User (POST /api/users/register)
func RegisterUser(c *gin.Context) {
  var payload struct {
    GameShiftID   string `json:"gameshiftId" binding:"required"`
    Email         string `json:"email" binding:"required,email"`
    Username      string `json:"username" binding:"required"`
    WalletType    string `json:"walletType" binding:"required"`
    WalletAddress string `json:"walletAddress" binding:"required"`
  }

  // Validate input
  if err := c.ShouldBindJSON(&payload); err != nil {
    log.Printf("❌ Invalid registration payload: %v", err)
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, all fields are required"})
    return
  }

  user := models.User{
    GameShiftID:   payload.GameShiftID,
    Email:         payload.Email,
    Username:      payload.Username,
    WalletType:    payload.WalletType,
    WalletAddress: payload.WalletAddress,
  }

  // Save to database
  if err := config.DB.Create(&user).Error; err != nil {
    log.Printf("❌ Registration failed: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login User (GET /api/users/login)
func LoginUser(c *gin.Context) {
  email := c.Query("email")

  if email == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
    return
  }

  var user models.User
  if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
    log.Printf("❌ Login failed, user not found: %s", email)
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  // Generate JWT token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":    user.ID,
    "email": user.Email,
    "exp":   time.Now().Add(24 * time.Hour).Unix(),
  })

  tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
  if err != nil {
    log.Printf("❌ Token generation failed: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "user":  user,
    "token": tokenString,
  })
}

// Get All Users (GET /api/users)
func GetAllUsers(c *gin.Context) {
  var users []models.User

  if err := config.DB.Find(&users).Error; err != nil {
    log.Printf("❌ Failed to retrieve users: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"users": users})
}

// Get User by Email (GET /api/users/email/:email)
func GetUserByEmail(c *gin.Context) {
  email := c.Param("email")

  var user models.User
  if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
    log.Printf("❌ User with email %s not found: %v", email, err)
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}

// Get User by Wallet Address (GET /api/users/wallet/:walletAddress)
func GetUserByWalletAddress(c *gin.Context) {
  walletAddress := c.Param("walletAddress")

  var user models.User
  if err := config.DB.Where("wallet_address = ?", walletAddress).First(&user).Error; err != nil {
    log.Printf("❌ User with wallet address %s not found: %v", walletAddress, err)
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update User by Email (PUT /api/users/email/:email)
func UpdateUserByEmail(c *gin.Context) {
  email := c.Param("email")

  var payload struct {
    GameShiftID   string `json:"gameshiftId"`
    Username      string `json:"username"`
    WalletType    string `json:"walletType"`
    WalletAddress string `json:"walletAddress"`
  }

  // Validate request payload
  if err := c.ShouldBindJSON(&payload); err != nil {
    log.Printf("❌ Invalid input for email %s: %v", email, err)
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  var user models.User

  // Check if user exists
  if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
    log.Printf("❌ User with email %s not found: %v", email, err)
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  // Update fields
  updatedFields := make(map[string]interface{})
  if payload.GameShiftID != "" {
    updatedFields["gameshiftId"] = payload.GameShiftID
  }
  if payload.Username != "" {
    updatedFields["username"] = payload.Username
  }
  if payload.WalletType != "" {
    updatedFields["walletType"] = payload.WalletType
  }
  if payload.WalletAddress != "" {
    updatedFields["walletAddress"] = payload.WalletAddress
  }

  // Save changes
  if err := config.DB.Model(&user).Updates(updatedFields).Error; err != nil {
    log.Printf("❌ Failed to update user: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

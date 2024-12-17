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

  // Validate request payload
  if err := c.ShouldBindJSON(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, all fields are required"})
    return
  }

  // Create new user object
  user := models.User{
    GameShiftID:   payload.GameShiftID,
    Email:         payload.Email,
    Username:      payload.Username,
    WalletType:    payload.WalletType,
    WalletAddress: payload.WalletAddress,
  }

  // Save user to database
  if err := config.DB.Create(&user).Error; err != nil {
    log.Printf("❌ Registration failed: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login User (POST /api/users/login)
func LoginUser(c *gin.Context) {
  var payload struct {
    Email string `json:"email" binding:"required,email"`
  }

  // Validate input payload
  if err := c.ShouldBindJSON(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
    return
  }

  // Look up user by email
  var user models.User
  if err := config.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  // Generate JWT token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":    user.ID,
    "email": user.Email,
    "exp":   time.Now().Add(24 * time.Hour).Unix(),
  })

  // Sign the token
  tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
  if err != nil {
    log.Printf("❌ Token generation failed: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    return
  }

  // Return the user and token
  c.JSON(http.StatusOK, gin.H{
    "user":  user,
    "token": tokenString,
  })
}

// Get All Users (GET /api/users)
func GetAllUsers(c *gin.Context) {
  var users []models.User

  // Retrieve all users from the database
  if err := config.DB.Find(&users).Error; err != nil {
    log.Printf("❌ Failed to retrieve users: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
    return
  }

  // Return the user list
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

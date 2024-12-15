package main

import (
  "log"
  "net/http"
  "os"
  "time"
  "fmt"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "github.com/golang-jwt/jwt/v5"
  "encoding/json"
)

// Global Database Instance
var DB *gorm.DB

// User Model
type User struct {
  ID            uint      `json:"id" gorm:"primaryKey"`
  GameShiftID   string    `json:"gameshiftId" gorm:"size:100;unique"`
  Username      string    `json:"username" gorm:"size:100"`
  Email         string    `json:"email" gorm:"size:255;unique"`
  WalletType    string    `json:"walletType" gorm:"size:100"`
  WalletAddress string    `json:"walletAddress" gorm:"size:255;unique"`
  CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

// Load Environment Variables
func init() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Println("Warning: .env file not found. Using system environment variables.")
  }
}

// Database Initialization
func ConnectDatabase() {
  dsn := os.Getenv("DATABASE_URL")
  if dsn == "" {
    dsn = "host=" + os.Getenv("DB_HOST") +
      " port=" + os.Getenv("DB_PORT") +
      " user=" + os.Getenv("DB_USERNAME") +
      " password=" + os.Getenv("DB_PASSWORD") +
      " dbname=" + os.Getenv("DB_DATABASE") +
      " sslmode=require"
  }

  var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Error), // Show errors only
  })
  if err != nil {
    log.Fatalf("Failed to connect to the database: %v", err)
  }

  // Add connection pooling
  sqlDB, _ := DB.DB()
  sqlDB.SetMaxIdleConns(10)
  sqlDB.SetMaxOpenConns(100)
  sqlDB.SetConnMaxLifetime(time.Hour)

  log.Println("Database connection established with SSL")
}

// Fetch Wallet Info from GameShift API
func fetchWalletInfo(gameShiftID string) (string, string, error) {
  apiKey := os.Getenv("GAMESHIFT_API_KEY")
  url := "https://api.gameshift.dev/nx/users/" + gameShiftID + "/wallet-address"

  client := &http.Client{Timeout: 10 * time.Second}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return "", "", err
  }
  req.Header.Set("Accept", "application/json")
  req.Header.Set("x-api-key", apiKey)

  resp, err := client.Do(req)
  if err != nil || resp.StatusCode != http.StatusOK {
    return "", "", err
  }
  defer resp.Body.Close()

  var walletData struct {
    Address       string `json:"address"`
    WalletProvider string `json:"walletProvider"`
  }
  if err := json.NewDecoder(resp.Body).Decode(&walletData); err != nil {
    return "", "", err
  }

  return walletData.Address, walletData.WalletProvider, nil
}

func main() {
  ConnectDatabase()

  // Initialize Gin Router
  r := gin.Default()

  // Root Route (Fix for GIN 404)
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to Rewarding Harvest API"})
  })

  // Health Check Route
  r.GET("/api/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "Server is running"})
  })

  // User Registration Route
  r.POST("/api/users/register", func(c *gin.Context) {
    var payload struct {
      GameShiftID string `json:"gameshiftId"`
      Email       string `json:"email"`
      Username    string `json:"username"`
    }

    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
      return
    }

    if payload.GameShiftID == "" || payload.Email == "" {
      c.JSON(http.StatusBadRequest, gin.H{
        "error": "Registration failed. Provide email and GameShift ID.",
      })
      return
    }

    walletAddress, walletType, err := fetchWalletInfo(payload.GameShiftID)
    if err != nil {
      log.Printf("Wallet Fetch Error: %v", err)
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Unable to retrieve wallet information.",
      })
      return
    }

    newUser := User{
      GameShiftID:   payload.GameShiftID,
      Email:         payload.Email,
      Username:      payload.Username,
      WalletType:    walletType,
      WalletAddress: walletAddress,
    }

    if err := DB.Create(&newUser).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Registration failed. Please contact support.",
      })
      return
    }

    c.JSON(http.StatusCreated, gin.H{
      "status":  "success",
      "message": "Registration successful! Welcome, " + payload.Username,
    })
  })

  // Login Route
  r.POST("/api/auth/login", func(c *gin.Context) {
    var payload struct {
      Email string `json:"email"`
    }

    if err := c.ShouldBindJSON(&payload); err != nil || payload.Email == "" {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Login failed. Email is required."})
      return
    }

    var user User
    if err := DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
      c.JSON(http.StatusNotFound, gin.H{
        "error": "No account found with that email address.",
      })
      return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "id":    user.ID,
      "email": user.Email,
      "exp":   time.Now().Add(time.Hour).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Login failed. Please try again later.",
      })
      return
    }

    c.JSON(http.StatusOK, gin.H{
      "status":  "success",
      "message": "Login successful!",
      "user":    user,
      "token":   tokenString,
    })
  })

  // Global 404 Handler
  r.NoRoute(func(c *gin.Context) {
    requestedPath := c.Request.URL.Path
    c.JSON(http.StatusNotFound, gin.H{
      "error": fmt.Sprintf("Resource not found at path: %s", requestedPath),
      "method": c.Request.Method,
    })
  })

  // Start the Server
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }
  r.Run(":" + port)
}

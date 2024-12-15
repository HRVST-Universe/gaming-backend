package config

import (
  "crypto/x509"
  "encoding/base64"
  "fmt"
  "log"
  "os"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
  // Load environment variables
  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")
  dbUser := os.Getenv("DB_USERNAME")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_DATABASE")
  sslMode := os.Getenv("SSL_MODE")
  caCertBase64 := os.Getenv("SSL_CA_CERT")

  // Ensure CA Certificate is loaded
  if caCertBase64 == "" {
    log.Fatal("Environment variable SSL_CA_CERT is not set.")
  }

  // Decode Base64-encoded CA Certificate
  caCert, err := base64.StdEncoding.DecodeString(caCertBase64)
  if err != nil {
    log.Fatalf("Failed to decode CA Certificate: %v", err)
  }

  // Create a certificate pool and add the CA Certificate
  certPool := x509.NewCertPool()
  if !certPool.AppendCertsFromPEM(caCert) {
    log.Fatal("Failed to append CA Certificate to Cert Pool")
  }

  // Create PostgreSQL DSN string
  dsn := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
  )

  // Connect to PostgreSQL using GORM
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
    log.Fatalf("Failed to connect to the database: %v", err)
  }

  fmt.Println("Database connection established with SSL")
}

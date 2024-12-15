package models

import "time"

type User struct {
  ID           uint      `gorm:"primaryKey"`
  Email        string    `gorm:"unique;size:255;not null"`
  Username     string    `gorm:"size:100"`
  WalletType   string    `gorm:"size:100"`
  WalletAddress string   `gorm:"unique;size:255"`
  CreatedAt    time.Time `gorm:"autoCreateTime"`
}

package models

import (
  "time"
)

type User struct {
  ID            uint      `gorm:"primaryKey" json:"id"`
  GameShiftID   string    `gorm:"column:gameshiftId;unique" json:"gameshiftId"`
  Email         string    `gorm:"column:email;unique" json:"email"`
  Username      string    `gorm:"column:username" json:"username"`
  WalletType    string    `gorm:"column:walletType" json:"walletType"`
  WalletAddress string    `gorm:"column:walletAddress" json:"walletAddress"`
  AvatarURL     string    `gorm:"column:avatarUrl" json:"avatarUrl"`
  CreatedAt     time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP" json:"createdAt"`
}

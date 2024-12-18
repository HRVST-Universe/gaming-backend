package models

import (
  "time"
)

type User struct {
  ID            uint      `gorm:"primaryKey;column:id" json:"id"`
  GameShiftID   string    `gorm:"column:gameshift_id;unique" json:"gameshiftId"`
  Email         string    `gorm:"column:email;unique" json:"email"`
  Username      string    `gorm:"column:username" json:"username"`
  WalletType    string    `gorm:"column:wallet_type" json:"walletType"`
  WalletAddress string    `gorm:"column:wallet_address;unique" json:"walletAddress"`
  AvatarURL     string    `gorm:"column:avatar_url" json:"avatarUrl"`
  CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

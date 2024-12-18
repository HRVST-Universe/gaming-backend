package models

import "time"

type Leaderboard struct {
  ID        uint      `gorm:"primaryKey;column:id" json:"id"`
  UserID    uint      `gorm:"column:user_id;notNull" json:"userId"`
  NFTID     uint      `gorm:"column:nft_id;notNull" json:"nftId"`
  Rank      int       `gorm:"column:rank" json:"rank"`
  XP        int       `gorm:"column:xp" json:"xp"`
  UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

package models

import "time"

type Leaderboard struct {
  ID        uint      `json:"id" gorm:"primaryKey"`
  UserID    uint      `json:"userId" gorm:"notNull"`
  NFTID     uint      `json:"nftId" gorm:"notNull"`
  Rank      int       `json:"rank"`
  XP        int       `json:"xp"`
  UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

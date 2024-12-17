package models

import "time"

type GameProgress struct {
  ID             uint      `json:"id" gorm:"primaryKey"`
  UserID         uint      `json:"userId" gorm:"notNull"`
  XP             int       `json:"xp" gorm:"default:0"`
  Level          int       `json:"level" gorm:"default:1"`
  CompletedTasks int       `json:"completedTasks" gorm:"default:0"`
  Badges         string    `json:"badges" gorm:"type:jsonb;default:'[]'"`
  LastActive     time.Time `json:"lastActive" gorm:"autoUpdateTime"`
}

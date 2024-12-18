package models

import "time"

type GameProgress struct {
  ID             uint      `gorm:"primaryKey;column:id" json:"id"`
  UserID         uint      `gorm:"column:user_id;notNull" json:"userId"`
  XP             int       `gorm:"column:xp;default:0" json:"xp"`
  Level          int       `gorm:"column:level;default:1" json:"level"`
  CompletedTasks int       `gorm:"column:completed_tasks;default:0" json:"completedTasks"`
  Badges         string    `gorm:"column:badges;type:jsonb;default:'[]'" json:"badges"`
  LastActive     time.Time `gorm:"column:last_active;autoUpdateTime" json:"lastActive"`
}

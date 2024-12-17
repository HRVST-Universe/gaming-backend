package models

import "time"

type PlayerAction struct {
  ID               uint      `json:"id" gorm:"primaryKey"`
  PlayerID         string    `json:"playerId" gorm:"size:100;notNull"`
  PlayerEmail      string    `json:"playerEmail" gorm:"size:255;notNull"`
  PlayerWallet     string    `json:"playerWallet" gorm:"size:255;notNull"`
  ActionType       string    `json:"actionType" gorm:"size:255;notNull"`
  ActionDescription string    `json:"actionDescription" gorm:"type:text"`
  Device           string    `json:"device" gorm:"size:255;default:'unknown'"`
  ActionTimestamp  time.Time `json:"actionTimestamp" gorm:"autoCreateTime"`
}

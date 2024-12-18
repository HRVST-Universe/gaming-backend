package models

import "time"

type PlayerAction struct {
  ID                uint      `gorm:"primaryKey;column:id" json:"id"`
  PlayerID          string    `gorm:"column:player_id;size:100;notNull" json:"playerId"`
  PlayerEmail       string    `gorm:"column:player_email;size:255;notNull" json:"playerEmail"`
  PlayerWallet      string    `gorm:"column:player_wallet;size:255;notNull" json:"playerWallet"`
  ActionType        string    `gorm:"column:action_type;size:255;notNull" json:"actionType"`
  ActionDescription string    `gorm:"column:action_description;type:text" json:"actionDescription"`
  Device            string    `gorm:"column:device;size:255;default:'unknown'" json:"device"`
  ActionTimestamp   time.Time `gorm:"column:action_timestamp;autoCreateTime" json:"actionTimestamp"`
}

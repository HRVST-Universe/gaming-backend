package models

import "time"

type NFTMetadata struct {
  ID           uint      `json:"id" gorm:"primaryKey"`
  UserID       uint      `json:"userId" gorm:"notNull"`
  Ownership    string    `json:"ownership" gorm:"size:100"`
  NFTID        string    `json:"nftId" gorm:"size:255;unique"`
  Squad        string    `json:"squad" gorm:"size:100"`
  XP           int       `json:"xp" gorm:"default:0"`
  Level        int       `json:"level" gorm:"default:1"`
  MetadataURI  string    `json:"metadataUri" gorm:"type:text"`
  UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

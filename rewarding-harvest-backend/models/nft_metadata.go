package models

import "time"

type NFTMetadata struct {
  ID           uint      `gorm:"primaryKey;column:id" json:"id"`
  UserID       uint      `gorm:"column:user_id;notNull" json:"userId"`
  Ownership    string    `gorm:"column:ownership;size:100" json:"ownership"`
  NFTID        string    `gorm:"column:nft_id;size:255;unique" json:"nftId"`
  Squad        string    `gorm:"column:squad;size:100" json:"squad"`
  XP           int       `gorm:"column:xp;default:0" json:"xp"`
  Level        int       `gorm:"column:level;default:1" json:"level"`
  MetadataURI  string    `gorm:"column:metadata_uri;type:text" json:"metadataUri"`
  UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

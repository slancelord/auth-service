package model

import "time"

type Session struct {
	ID           string `gorm:"primaryKey;autoIncrement"`
	UserID       string `gorm:"type:uuid"`
	RefreshToken string
	TokenPairId  string `gorm:"type:uuid"`
	UserAgent    string
	IPAddress    string
	CreatedAt    time.Time
	ExpiresAt    time.Time
}

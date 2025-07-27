package model

import "time"

type Session struct {
	ID           uint
	GUID         string
	RefreshToken string
	CreatedAt    time.Time
	ExpiresAt    time.Time
}

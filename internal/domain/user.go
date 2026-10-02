package domain

import "time"

type User struct {
	ID           string
	Username     string
	PasswordHash string
	Avatar       int
	Clicks       int64
	CreatedAt    time.Time
}

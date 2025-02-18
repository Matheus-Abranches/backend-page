package model

import (
	"time"
)

type SentEmail struct {
	ID      uint      `json:"id"`
	Date    time.Time `json:"date"`
	Version string    `json:"version"`
	Product string    `json:"product"`
	UserID  uint      `json:"user_id"`
}

package model

import "time"

type User struct {
	ID                uint      `json:"id"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	AgreeTerms        bool      `json:"agree_terms"`
	AgreeReceivesNews bool      `json:"agree_receives_news"`
	CreatAt           time.Time `json:"created_at"`
	DeleteAt          time.Time `json:"deleted_at"`
}

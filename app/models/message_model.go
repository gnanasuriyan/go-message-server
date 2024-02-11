package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey" db:"id"`
	FkUser    uint      `json:"fk_user" db:"fk_user"`
	Content   string    `json:"content" db:"content"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type MessageCreate struct {
	FkUser  uint   `json:"fk_user"`
	Content string `json:"content"`
}

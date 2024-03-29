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

type MessageCreateDto struct {
	Content string `json:"content"`
}

type MessageResponseDto struct {
	ID                    uint      `json:"id"`
	Content               string    `json:"content"`
	PostedBy              string    `json:"posted_by"`
	CreatedAt             time.Time `json:"created_at"`
	IsPostedByCurrentUser bool      `json:"is_posted_by_current_user"`
}

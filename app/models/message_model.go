package models

import "time"

type Message struct {
	ID        int       `json:"id"`
	FkUser    int       `json:"fk_user"`
	Content   string    `json:"content"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

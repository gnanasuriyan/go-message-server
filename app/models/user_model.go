package models

import "time"

type User struct {
	ID        int       `json:"id" migrations:"id"`
	Username  string    `json:"username" migrations:"username"`
	Password  string    `json:"password" migrations:"password"`
	Active    bool      `json:"active" migrations:"active"`
	CreatedAt time.Time `json:"created_at" migrations:"created_at"`
	UpdatedAt time.Time `json:"updated_at" migrations:"updated_at"`
}

type UserCreate struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

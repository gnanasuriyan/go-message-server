package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey" db:"id"`
	Username     string    `json:"username" db:"username"`
	Password     string    `json:"password" db:"password"`
	Active       bool      `json:"active" db:"active"`
	ShadowActive bool      `json:"shadow_active" db:"shadow_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type UserCreate struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

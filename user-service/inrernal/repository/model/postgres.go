package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(150);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}


type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
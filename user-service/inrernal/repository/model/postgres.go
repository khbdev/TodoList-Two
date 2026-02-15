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

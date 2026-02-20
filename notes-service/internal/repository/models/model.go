package models

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"` 
	Title     string    `gorm:"type:varchar(255);not null"`
	Text      string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
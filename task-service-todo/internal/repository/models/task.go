package models

import "time"

type Reminder struct {
	ID        uint      `gorm:"primaryKey"` 

	UserID    uint      `gorm:"not null;index"` 

	Task      string    `gorm:"type:text;not null"`

	RemindAt  time.Time `gorm:"type:timestamptz;not null;index:idx_notified_remind"` 

	Notified  bool      `gorm:"default:false;index:idx_notified_remind"` 

	CreatedAt time.Time
	UpdatedAt time.Time
}
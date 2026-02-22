package domain

import (
	"context"
	"time"

	"task-service/internal/repository/models"
)


type ReminderUsecase interface {
	Create(ctx context.Context, userID uint, task string, remindAt time.Time) (*models.Reminder, error)

	GetByUser(ctx context.Context, userID uint) ([]models.Reminder, error)

	GetByID(ctx context.Context, userID uint, reminderID uint) (*models.Reminder, error)

	Update(ctx context.Context, userID uint, reminderID uint, task string, remindAt time.Time, notified bool) (*models.Reminder, error)
	
	Delete(ctx context.Context, userID uint, reminderID uint) error
}
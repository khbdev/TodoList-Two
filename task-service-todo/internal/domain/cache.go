package domain



import (
	"context"
	"time"
	"task-service/internal/repository/models"
)

type ReminderCache interface {

	GetListByUser(ctx context.Context, userID uint) ([]models.Reminder, error)
	SetListByUser(ctx context.Context, userID uint, reminders []models.Reminder, ttl time.Duration) error
	DeleteListByUser(ctx context.Context, userID uint) error

	GetByID(ctx context.Context, userID uint, reminderID uint) (*models.Reminder, error)
	SetByID(ctx context.Context, reminder *models.Reminder, ttl time.Duration) error
	DeleteByID(ctx context.Context, userID uint, reminderID uint) error
}
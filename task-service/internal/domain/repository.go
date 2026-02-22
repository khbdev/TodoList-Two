package domain

import (
	"context"
	"task-service/internal/repository/models"
)

type TaskService interface {
	Create(ctx context.Context, task *models.Reminder) (*models.Reminder, error)

	GetByID(ctx context.Context, taskID int64) (*models.Reminder, error) 

	GetByUser(ctx context.Context, userID int64) ([]models.Reminder, error)

	Update(ctx context.Context, task *models.Reminder) (*models.Reminder, error)

	Delete(ctx context.Context, taskID int64) error
}
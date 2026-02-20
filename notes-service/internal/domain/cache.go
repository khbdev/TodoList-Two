package domain

import (
	"context"
	"time"

	"notes-service/internal/repository/models"
)

type TodoCache interface {

	GetListByUser(ctx context.Context, userID uint) ([]models.Todo, error)
	SetListByUser(ctx context.Context, userID uint, todos []models.Todo, ttl time.Duration) error
	DeleteListByUser(ctx context.Context, userID uint) error

	
	GetByID(ctx context.Context, userID uint, todoID uint) (*models.Todo, error)
	SetByID(ctx context.Context, todo *models.Todo, ttl time.Duration) error
	DeleteByID(ctx context.Context, userID uint, todoID uint) error
}
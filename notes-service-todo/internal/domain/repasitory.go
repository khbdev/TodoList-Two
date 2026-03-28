package domain

import (
	"context"
	"notes-service/internal/repository/models"
)

type TodoRepository interface {
	Create(ctx context.Context, t *models.Todo) error


	GetAll(ctx context.Context, userID uint) ([]models.Todo, error)


	GetByID(ctx context.Context, userID uint, todoID uint) (*models.Todo, error)


	Update(ctx context.Context, t *models.Todo) error

	
	Delete(ctx context.Context, userID uint, todoID uint) error
}
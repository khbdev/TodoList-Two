package postgres

import (
	"context"
	"errors"
	"notes-service/internal/repository/models"

	"gorm.io/gorm"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) Create(ctx context.Context, t *models.Todo) error {
	return r.db.WithContext(ctx).Create(t).Error
}


func (r *TodoRepo) GetByID(ctx context.Context, userID uint, todoID uint) (*models.Todo, error) {
	var t models.Todo

	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", todoID, userID).
		First(&t).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}

	return &t, nil
}


func (r *TodoRepo) GetAll(ctx context.Context, userID uint) ([]models.Todo, error) {
	var todos []models.Todo

	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}


func (r *TodoRepo) Update(ctx context.Context, t *models.Todo) error {

	tx := r.db.WithContext(ctx).
		Model(&models.Todo{}).
		Where("id = ? AND user_id = ?", t.ID, t.UserID).
		Updates(map[string]any{
			"title": t.Title,
			"text":  t.Text,
		})

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return ErrTodoNotFound
	}
	return nil
}


func (r *TodoRepo) Delete(ctx context.Context, userID uint, todoID uint) error {
	tx := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", todoID, userID).
		Delete(&models.Todo{})

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return ErrTodoNotFound
	}
	return nil
}
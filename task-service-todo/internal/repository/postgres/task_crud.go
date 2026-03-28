package postgres

import (
	"context"
	"errors"
	"task-service/internal/domain"
	"task-service/internal/repository/models"

	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}


func NewTaskRepo(db *gorm.DB) domain.TaskService {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) Create(ctx context.Context, task *models.Reminder) (*models.Reminder, error) {
	if err := r.db.WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepo) GetByID(ctx context.Context, taskID int64) (*models.Reminder, error) {
	var task models.Reminder

	err := r.db.WithContext(ctx).
		First(&task, taskID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) GetByUser(ctx context.Context, userID int64) ([]models.Reminder, error) {
	var tasks []models.Reminder

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepo) Update(ctx context.Context, task *models.Reminder) (*models.Reminder, error) {
	err := r.db.WithContext(ctx).
		Save(task).Error

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepo) Delete(ctx context.Context, taskID int64) error {
	result := r.db.WithContext(ctx).
		Delete(&models.Reminder{}, taskID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
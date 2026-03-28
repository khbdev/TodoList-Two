package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"task-service/internal/domain"
	"task-service/internal/repository/models"
)

var (
	ErrInvalidTask     = errors.New("task is required")
	ErrInvalidUser     = errors.New("user_id is required")
	ErrInvalidReminder = errors.New("reminder_id is required")
	ErrInvalidRemindAt = errors.New("remind_at is required")
)

type ReminderService struct {
	repo  domain.TaskService     
	cache domain.ReminderCache   
	ttl   time.Duration
}

func NewReminderService(repo domain.TaskService, cache domain.ReminderCache, ttl time.Duration) *ReminderService {
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	return &ReminderService{
		repo:  repo,
		cache: cache,
		ttl:   ttl,
	}
}


func (s *ReminderService) Create(ctx context.Context, userID uint, task string, remindAt time.Time) (*models.Reminder, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if strings.TrimSpace(task) == "" {
		return nil, ErrInvalidTask
	}
	if remindAt.IsZero() {
		return nil, ErrInvalidRemindAt
	}

	r := &models.Reminder{
		UserID:   userID,
		Task:     strings.TrimSpace(task),
		RemindAt: remindAt.UTC(),
		Notified: false,
	}

	created, err := s.repo.Create(ctx, r)
	if err != nil {
		return nil, err
	}


	if s.cache != nil && created != nil && created.ID != 0 {
		_ = s.cache.DeleteListByUser(ctx, userID)
		_ = s.cache.SetByID(ctx, created, s.ttl)
	}

	return created, nil
}


func (s *ReminderService) GetByUser(ctx context.Context, userID uint) ([]models.Reminder, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}


	if s.cache != nil {
		cached, err := s.cache.GetListByUser(ctx, userID)
		if err != nil {
			return nil, err
		}
		if cached != nil {
			return cached, nil
		}
	}

	reminders, err := s.repo.GetByUser(ctx, int64(userID))
	if err != nil {
		return nil, err
	}

	if s.cache != nil {
		_ = s.cache.SetListByUser(ctx, userID, reminders, s.ttl)
	}

	return reminders, nil
}


func (s *ReminderService) GetByID(ctx context.Context, userID uint, reminderID uint) (*models.Reminder, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if reminderID == 0 {
		return nil, ErrInvalidReminder
	}


	if s.cache != nil {
		cached, err := s.cache.GetByID(ctx, userID, reminderID)
		if err != nil {
			return nil, err
		}
		if cached != nil {
			return cached, nil
		}
	}

	rem, err := s.repo.GetByID(ctx, int64(reminderID))
	if err != nil {
		return nil, err
	}


	if rem != nil && rem.UserID != userID {
		return nil, nil
	}

	if s.cache != nil && rem != nil {
		_ = s.cache.SetByID(ctx, rem, s.ttl)
	}

	return rem, nil
}


func (s *ReminderService) Update(ctx context.Context, userID uint, reminderID uint, task string, remindAt time.Time, notified bool) (*models.Reminder, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if reminderID == 0 {
		return nil, ErrInvalidReminder
	}
	if strings.TrimSpace(task) == "" {
		return nil, ErrInvalidTask
	}
	if remindAt.IsZero() {
		return nil, ErrInvalidRemindAt
	}

	existing, err := s.repo.GetByID(ctx, int64(reminderID))
	if err != nil {
		return nil, err
	}
	if existing == nil || existing.UserID != userID {
		return nil, nil
	}

	existing.Task = strings.TrimSpace(task)
	existing.RemindAt = remindAt.UTC()
	existing.Notified = notified

	updated, err := s.repo.Update(ctx, existing)
	if err != nil {
		return nil, err
	}


	if s.cache != nil {
		_ = s.cache.DeleteByID(ctx, userID, reminderID)
		_ = s.cache.DeleteListByUser(ctx, userID)
		_ = s.cache.SetByID(ctx, updated, s.ttl)
	}

	return updated, nil
}


func (s *ReminderService) Delete(ctx context.Context, userID uint, reminderID uint) error {
	if userID == 0 {
		return ErrInvalidUser
	}
	if reminderID == 0 {
		return ErrInvalidReminder
	}

	existing, err := s.repo.GetByID(ctx, int64(reminderID))
	if err != nil {
		return err
	}
	if existing == nil || existing.UserID != userID {
		return nil
	}

	if err := s.repo.Delete(ctx, int64(reminderID)); err != nil {
		return err
	}

	if s.cache != nil {
		_ = s.cache.DeleteByID(ctx, userID, reminderID)
		_ = s.cache.DeleteListByUser(ctx, userID)
	}

	return nil
}
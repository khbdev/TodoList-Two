package usecase

import (
	"context"
	"errors"
	"notes-service/internal/domain"
	"notes-service/internal/repository/models"

	"strings"
)

var (
	ErrInvalidTitle = errors.New("title is required")
	ErrInvalidUser  = errors.New("user_id is required")
	ErrInvalidTodo  = errors.New("todo_id is required")
)

type TodoService struct {
	repo domain.TodoRepository
}

func NewTodoService(repo domain.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}


func (s *TodoService) Create(ctx context.Context, userID uint, title string, text string) (*models.Todo, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if strings.TrimSpace(title) == "" {
		return nil, ErrInvalidTitle
	}

	t := &models.Todo{
		UserID: userID,
		Title:  strings.TrimSpace(title),
		Text:   strings.TrimSpace(text),
	}

	if err := s.repo.Create(ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TodoService) GetAll(ctx context.Context, userID uint) ([]models.Todo, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	return s.repo.GetAll(ctx, userID)
}

func (s *TodoService) GetByID(ctx context.Context, userID uint, todoID uint) (*models.Todo, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if todoID == 0 {
		return nil, ErrInvalidTodo
	}
	return s.repo.GetByID(ctx, userID, todoID)
}


func (s *TodoService) Update(ctx context.Context, userID uint, todoID uint, title string, text string) (*models.Todo, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if todoID == 0 {
		return nil, ErrInvalidTodo
	}
	if strings.TrimSpace(title) == "" {
		return nil, ErrInvalidTitle
	}

	
	existing, err := s.repo.GetByID(ctx, userID, todoID)
	if err != nil {
		return nil, err
	}

	existing.Title = strings.TrimSpace(title)
	existing.Text = strings.TrimSpace(text)

	if err := s.repo.Update(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *TodoService) Delete(ctx context.Context, userID uint, todoID uint) error {
	if userID == 0 {
		return ErrInvalidUser
	}
	if todoID == 0 {
		return ErrInvalidTodo
	}

	_, err := s.repo.GetByID(ctx, userID, todoID)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, userID, todoID)
}
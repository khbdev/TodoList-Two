package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"notes-service/internal/domain"
	"notes-service/internal/repository/models"
)

var (
	ErrInvalidTitle = errors.New("title is required")
	ErrInvalidUser  = errors.New("user_id is required")
	ErrInvalidTodo  = errors.New("todo_id is required")
)

type TodoService struct {
	repo  domain.TodoRepository
	cache domain.TodoCache
	ttl   time.Duration
}

func NewTodoService(repo domain.TodoRepository, cache domain.TodoCache, ttl time.Duration) *TodoService {
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	return &TodoService{
		repo:  repo,
		cache: cache,
		ttl:   ttl,
	}
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

	
	if s.cache != nil && t.ID != 0 {
		_ = s.cache.DeleteListByUser(ctx, userID) 
		_ = s.cache.SetByID(ctx, t, s.ttl)      
	}

	return t, nil
}

func (s *TodoService) GetAll(ctx context.Context, userID uint) ([]models.Todo, error) {
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

	todos, err := s.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	if s.cache != nil {
		_ = s.cache.SetListByUser(ctx, userID, todos, s.ttl)
	}

	return todos, nil
}

func (s *TodoService) GetByID(ctx context.Context, userID uint, todoID uint) (*models.Todo, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if todoID == 0 {
		return nil, ErrInvalidTodo
	}

	
	if s.cache != nil {
		cached, err := s.cache.GetByID(ctx, userID, todoID)
		if err != nil {
			return nil, err
		}
		if cached != nil {
			return cached, nil
		}
	}

	todo, err := s.repo.GetByID(ctx, userID, todoID)
	if err != nil {
		return nil, err
	}

	if s.cache != nil && todo != nil {
		_ = s.cache.SetByID(ctx, todo, s.ttl)
	}

	return todo, nil
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


	if s.cache != nil {
		_ = s.cache.DeleteByID(ctx, userID, todoID)   
		_ = s.cache.DeleteListByUser(ctx, userID)   
		_ = s.cache.SetByID(ctx, existing, s.ttl)    
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

	if err := s.repo.Delete(ctx, userID, todoID); err != nil {
		return err
	}


	if s.cache != nil {
		_ = s.cache.DeleteByID(ctx, userID, todoID)
		_ = s.cache.DeleteListByUser(ctx, userID)
	}

	return nil
}
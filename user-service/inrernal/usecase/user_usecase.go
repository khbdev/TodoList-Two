package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"user-service/inrernal/domain"
	"user-service/inrernal/repository/model"
	"user-service/pkg"
)

type UserService struct {
	repo  domain.UserRepository
	cache domain.UserCache
	ttl   time.Duration
}

func NewUserService(repo domain.UserRepository, cache domain.UserCache, ttl time.Duration) *UserService {
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	return &UserService{
		repo:  repo,
		cache: cache,
		ttl:   ttl,
	}
}

func (s *UserService) Create(ctx context.Context, name string, email string, password string) (*model.User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name required")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email required")
	}
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password required")
	}

	hashed, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:         strings.TrimSpace(name),
		Email:        strings.ToLower(strings.TrimSpace(email)),
		PasswordHash: hashed,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	// write-through
	if s.cache != nil && user.ID != 0 {
		_ = s.cache.SetByID(ctx, user, s.ttl)
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, id int, name string, email string) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	user, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(name) != "" {
		user.Name = strings.TrimSpace(name)
	}
	if strings.TrimSpace(email) != "" {
		user.Email = strings.ToLower(strings.TrimSpace(email))
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	// write-through
	if s.cache != nil {
		_ = s.cache.DeleteByID(ctx, id)
		_ = s.cache.SetByID(ctx, user, s.ttl)
	}

	return user, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	// read-through
	if s.cache != nil {
		cached, err := s.cache.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if cached != nil {
			return cached, nil
		}
	}

	user, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	if s.cache != nil && user != nil {
		_ = s.cache.SetByID(ctx, user, s.ttl)
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	if err := s.repo.Delete(ctx, uint(id)); err != nil {
		return err
	}

	// delete-through
	if s.cache != nil {
		_ = s.cache.DeleteByID(ctx, id)
	}

	return nil
}

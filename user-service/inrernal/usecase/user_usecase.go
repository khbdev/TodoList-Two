package usecase

import (
	"context"
	"errors"
	"strings"

	"user-service/inrernal/domain"
	"user-service/inrernal/repository/model"
	"user-service/pkg"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, name string, password string) (*model.User, error) {

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name required")
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
		PasswordHash: hashed,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
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

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
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

	return s.repo.GetByID(ctx, uint(id))
}

func (s *UserService) Delete(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.repo.Delete(ctx, uint(id))
}

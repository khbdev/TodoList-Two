package postgres

import (
	"context"
	"errors"
	"user-service/inrernal/repository/model"

	"gorm.io/gorm"
)


var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists  = errors.New("email already exists")
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, u *model.User) error {
	var existing model.User


	err := r.db.WithContext(ctx).
		Where("email = ?", u.Email).
		First(&existing).Error

	if err == nil {
		return ErrEmailExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return r.db.WithContext(ctx).Create(u).Error
}



func (r *UserRepo) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var u model.User
	err := r.db.WithContext(ctx).First(&u, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}


func (r *UserRepo) Update(ctx context.Context, u *model.User) error {
	tx := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", u.ID).Updates(u)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}


func (r *UserRepo) Delete(ctx context.Context, id uint) error {
	tx := r.db.WithContext(ctx).Delete(&model.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

package domain

import "user-service/internal/repositroy/model"


type UserRepository interface {
	Create(user model.User) (error)
	
}
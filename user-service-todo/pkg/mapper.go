package pkg

import "user-service/inrernal/repository/model"


func ToUserResponse(u *model.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserResponseList(users []model.User) []model.UserResponse {
	res := make([]model.UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, model.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return res
}
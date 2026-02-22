package pkg

import (
	"errors"

	"task-service/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, usecase.ErrInvalidTask),
		errors.Is(err, usecase.ErrInvalidUser),
		errors.Is(err, usecase.ErrInvalidReminder),
		errors.Is(err, usecase.ErrInvalidRemindAt):
		return status.Error(codes.InvalidArgument, err.Error())

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
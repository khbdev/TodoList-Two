package pkg

import (
	

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(err error) error {
	if err == nil {
		return nil
	}
	switch err.Error() {
	case "title is required", "user_id is required", "todo_id is required":
		return status.Error(codes.InvalidArgument, err.Error())
	}	
	return status.Error(codes.Internal, err.Error())
}
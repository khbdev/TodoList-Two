package pkg

import (
	"task-service/internal/repository/models"
	"time"

	taskpb "github.com/khbdev/todolist-proto/proto/task"
)

func ToProtoNote(t *models.Reminder) *taskpb.Task {
	if t == nil {
		return nil
	}
	return &taskpb.Task{
		Id:        int64(t.ID),
		UserId:    int64(t.UserID),
		Title:     t.Task,
		RemindAt: t.RemindAt.Format(time.RFC3339),
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
		UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
	}
}

func ToProtoNotes(list []models.Reminder) []*taskpb.Task {
	out := make([]*taskpb.Task, 0, len(list))
	for i := range list {
		item := list[i] 
		out = append(out, ToProtoNote(&item))
	}
	return out
}
package pkg

import (
	"time"

	"task-service/internal/repository/models"

	taskpb "github.com/khbdev/todolist-proto/proto/task"
)

func ToProtoTask(t *models.Reminder) *taskpb.Task {
	if t == nil {
		return nil
	}

	return &taskpb.Task{
		Id:        int64(t.ID),
		UserId:    int64(t.UserID),
		Title:     t.Task,
		RemindAt:  t.RemindAt.UTC().Format(time.RFC3339),
		CreatedAt: t.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: t.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func ToProtoTasks(list []models.Reminder) []*taskpb.Task {
	out := make([]*taskpb.Task, 0, len(list))
	for i := range list {
		item := list[i]
		out = append(out, ToProtoTask(&item))
	}
	return out
}
package pkg

import (
	"time"

	"notes-service/internal/repository/models"

	notespb "github.com/khbdev/todolist-proto/proto/notes"
)

func ToProtoNote(t *models.Todo) *notespb.Note {
	if t == nil {
		return nil
	}
	return &notespb.Note{
		Id:        int64(t.ID),
		UserId:    int64(t.UserID),
		Title:     t.Title,
		Text:      t.Text,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
		UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
	}
}

func ToProtoNotes(list []models.Todo) []*notespb.Note {
	out := make([]*notespb.Note, 0, len(list))
	for i := range list {
		item := list[i] 
		out = append(out, ToProtoNote(&item))
	}
	return out
}
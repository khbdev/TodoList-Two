package worker

import (
	"context"
	"task-service/internal/event"
	"task-service/internal/repository/models"
	"time"

	"gorm.io/gorm"
)



type ReminderWorker  struct {
	db *gorm.DB
	prod *event.Producer
}



func NewReminderWorker(db *gorm.DB, prod *event.Producer) *ReminderWorker {
	return &ReminderWorker{
		db:   db,
		prod: prod,
	}
}


func (w *ReminderWorker) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
			case <-time.After(5 * time.Second):
					var list []models.Reminder
		}
	}
}
package worker

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"

	"task-service/internal/event"
	"task-service/internal/repository/models"
)

type ReminderWorker struct {
	db   *gorm.DB
	prod *event.Producer
}

func NewReminderWorker(db *gorm.DB, prod *event.Producer) *ReminderWorker {
	return &ReminderWorker{db: db, prod: prod}
}

func (w *ReminderWorker) Run(ctx context.Context) {
	log.Println("reminder worker started")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("reminder worker stopped:", ctx.Err())
			return

		case <-ticker.C:
			var list []models.Reminder

			// ✅ DB o‘zi NOW() bilan hisoblaydi (timezone muammosi yo‘q)
			if err := w.db.WithContext(ctx).
				Where("notified = false AND remind_at <= NOW()").
				Find(&list).Error; err != nil {
				log.Println("worker find error:", err)
				continue
			}

			if len(list) == 0 {
				// log.Println("worker: no due reminders")
				continue
			}

			for _, r := range list {
				// 1) publish
				if err := w.prod.PublishTask(ctx, int64(r.UserID), r.Task); err != nil {
					log.Println("publish error reminder_id=", r.ID, "err=", err)
					continue
				}

				// 2) notified=true (faqat hali false bo‘lsa)
				if err := w.db.WithContext(ctx).
					Model(&models.Reminder{}).
					Where("id = ? AND notified = false", r.ID).
					Update("notified", true).Error; err != nil {
					log.Println("update notified error reminder_id=", r.ID, "err=", err)
					continue
				}

				log.Println("reminder sent + marked notified id=", r.ID)
			}
		}
	}
}
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"task-service/internal/domain"
	"task-service/internal/repository/models"

	"github.com/redis/go-redis/v9"
)

type ReminderCache struct {
	rdb *redis.Client
}

func NewReminderCache(rdb *redis.Client) domain.ReminderCache {
	return &ReminderCache{rdb: rdb}
}


func listKey(userID uint) string {
	return fmt.Sprintf("reminders:%d", userID)
}

func itemKey(userID uint, reminderID uint) string {
	return fmt.Sprintf("reminders:item:%d:%d", userID, reminderID)
}



func (c *ReminderCache) GetListByUser(ctx context.Context, userID uint) ([]models.Reminder, error) {
	val, err := c.rdb.Get(ctx, listKey(userID)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var reminders []models.Reminder
	if err := json.Unmarshal([]byte(val), &reminders); err != nil {
		return nil, err
	}

	return reminders, nil
}

func (c *ReminderCache) SetListByUser(ctx context.Context, userID uint, reminders []models.Reminder, ttl time.Duration) error {
	data, err := json.Marshal(reminders)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, listKey(userID), data, ttl).Err()
}

func (c *ReminderCache) DeleteListByUser(ctx context.Context, userID uint) error {
	return c.rdb.Del(ctx, listKey(userID)).Err()
}



func (c *ReminderCache) GetByID(ctx context.Context, userID uint, reminderID uint) (*models.Reminder, error) {
	val, err := c.rdb.Get(ctx, itemKey(userID, reminderID)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var reminder models.Reminder
	if err := json.Unmarshal([]byte(val), &reminder); err != nil {
		return nil, err
	}

	return &reminder, nil
}

func (c *ReminderCache) SetByID(ctx context.Context, reminder *models.Reminder, ttl time.Duration) error {
	if reminder == nil {
		return nil
	}

	data, err := json.Marshal(reminder)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, itemKey(reminder.UserID, reminder.ID), data, ttl).Err()
}

func (c *ReminderCache) DeleteByID(ctx context.Context, userID uint, reminderID uint) error {
	return c.rdb.Del(ctx, itemKey(userID, reminderID)).Err()
}
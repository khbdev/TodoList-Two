package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"notes-service/internal/domain"
	"notes-service/internal/repository/models"

	"github.com/redis/go-redis/v9"
)

type TodoCache struct {
	rdb *redis.Client
}

func NewTodoCache(rdb *redis.Client) domain.TodoCache {
	return &TodoCache{rdb: rdb}
}


func listKey(userID uint) string {
	return fmt.Sprintf("notes:%d", userID)
}

func itemKey(userID uint, noteID uint) string {
	return fmt.Sprintf("notes:item:%d:%d", userID, noteID)
}



func (c *TodoCache) GetListByUser(ctx context.Context, userID uint) ([]models.Todo, error) {
	val, err := c.rdb.Get(ctx, listKey(userID)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	if err := json.Unmarshal([]byte(val), &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (c *TodoCache) SetListByUser(ctx context.Context, userID uint, todos []models.Todo, ttl time.Duration) error {

	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, listKey(userID), data, ttl).Err()
}

func (c *TodoCache) DeleteListByUser(ctx context.Context, userID uint) error {
	return c.rdb.Del(ctx, listKey(userID)).Err()
}



func (c *TodoCache) GetByID(ctx context.Context, userID uint, todoID uint) (*models.Todo, error) {
	val, err := c.rdb.Get(ctx, itemKey(userID, todoID)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var todo models.Todo
	if err := json.Unmarshal([]byte(val), &todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (c *TodoCache) SetByID(ctx context.Context, todo *models.Todo, ttl time.Duration) error {
	if todo == nil {
		return nil
	}

	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, itemKey(todo.UserID, todo.ID), data, ttl).Err()
}

func (c *TodoCache) DeleteByID(ctx context.Context, userID uint, todoID uint) error {
	return c.rdb.Del(ctx, itemKey(userID, todoID)).Err()
}
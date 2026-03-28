package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"user-service/inrernal/domain"
	"user-service/inrernal/repository/model"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	rdb *redis.Client
}

func NewUserCache(rdb *redis.Client) domain.UserCache {
	return &UserCache{rdb: rdb}
}

func key(id int) string {
	return fmt.Sprintf("user:%d", id)
}

func (c *UserCache) GetByID(ctx context.Context, id int) (*model.User, error) {
	val, err := c.rdb.Get(ctx, key(id)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserCache) SetByID(ctx context.Context, user *model.User, ttl time.Duration) error {
	if user == nil {
		return nil
	}

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key(int(user.ID)), data, ttl).Err()
}

func (c *UserCache) DeleteByID(ctx context.Context, id int) error {
	return c.rdb.Del(ctx, key(id)).Err()
}

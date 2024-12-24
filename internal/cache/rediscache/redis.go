package rediscache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
	ttl    time.Duration
}

func New(host string, port int, password string, db int, ttl time.Duration) Redis {
	return Redis{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: password,
			DB:       db,
		}),
		ttl: ttl,
	}
}

func (s Redis) Get(ctx context.Context, key string) (string, bool, error) {
	res, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return ``, false, nil
		} else {
			return ``, false, err
        }
	}

	return res, true, nil
}

func (s Redis) Set(ctx context.Context, key, value string) error {
	return s.client.Set(ctx, key, value, s.ttl).Err()
}

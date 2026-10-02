package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func New(cfg Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Redis{
		Client: client,
	}
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

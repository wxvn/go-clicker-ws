package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	redisClient "github.com/wxvn/go-clicker-ws/internal/redis"
)

var ErrSessionNotFound = errors.New("session not found")

type Store interface {
	CreateSession(ctx context.Context, userID string) (string, error)
	GetSession(ctx context.Context, sessionID string) (string, error)
	RefreshSession(ctx context.Context, sessionID string) error
	DeleteSession(ctx context.Context, sessionID string) error
}

type RedisStore struct {
	client *redisClient.Redis
	ttl    time.Duration
}

func NewStore(client *redisClient.Redis, ttl time.Duration) *RedisStore {
	return &RedisStore{
		client: client,
		ttl:    ttl,
	}
}

func (s *RedisStore) CreateSession(ctx context.Context, userID string) (string, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("session:%s", sessionID)

	err = s.client.Set(ctx, key, userID, s.ttl).Err()
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func (s *RedisStore) GetSession(ctx context.Context, sessionID string) (string, error) {
	key := fmt.Sprintf("session:%s", sessionID)

	userID, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrSessionNotFound
		}

		return "", err
	}

	return userID, nil
}

func (s *RedisStore) RefreshSession(ctx context.Context, sessionID string) error {
	key := fmt.Sprintf("session:%s", sessionID)

	ok, err := s.client.Expire(ctx, key, s.ttl).Result()
	if err != nil {
		return err
	}

	if !ok {
		return ErrSessionNotFound
	}

	return nil
}

func (s *RedisStore) DeleteSession(ctx context.Context, sessionID string) error {
	key := fmt.Sprintf("session:%s", sessionID)

	return s.client.Del(ctx, key).Err()
}

func generateSessionID() (string, error) {
	data := make([]byte, 32)

	if _, err := rand.Read(data); err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

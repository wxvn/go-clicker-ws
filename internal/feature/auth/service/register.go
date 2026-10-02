package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/wxvn/go-clicker-ws/internal/domain"
)

func (s *AuthService) Register(ctx context.Context, auth domain.Auth) (domain.User, string, error) {
	passwordHash, err := s.hasher.Hash(auth.Password)
	if err != nil {
		return domain.User{}, "", fmt.Errorf("hash password: %w", err)
	}

	createUser := &domain.User{
		Username:     auth.Username,
		PasswordHash: passwordHash,
		Avatar:       rand.Intn(6) + 1,
		Clicks:       0,
		CreatedAt:    time.Now(),
	}

	user, err := s.userRepository.CreateUser(ctx, createUser)
	if err != nil {
		return domain.User{}, "", fmt.Errorf("create user: %w", err)
	}

	sessionID, err := s.sessionStore.CreateSession(ctx, user.ID)
	if err != nil {
		return domain.User{}, "", fmt.Errorf("create session: %w", err)
	}

	return *user, sessionID, nil
}

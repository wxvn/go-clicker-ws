package service

import (
	"context"
	"fmt"

	"github.com/wxvn/go-clicker-ws/internal/domain"
)

func (s *AuthService) Login(ctx context.Context, auth domain.Auth) (domain.User, string, error) {

	user, err := s.userRepository.GetUserByUsername(ctx, auth.Username)
	if err != nil {
		return domain.User{}, "", fmt.Errorf("get user: %w", err)
	}

	if err := s.hasher.Compare(auth.Password, user.PasswordHash); err != nil {
		return domain.User{}, "", fmt.Errorf("compair password: %w", err)
	}

	sessionID, err := s.sessionStore.CreateSession(ctx, user.ID)
	if err != nil {
		return domain.User{}, "", fmt.Errorf("create session: %w", err)
	}

	return *user, sessionID, nil
}

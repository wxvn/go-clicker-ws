package service

import (
	"context"
	"fmt"
)

func (s *AuthService) Logout(ctx context.Context, sessionID string) error {

	if err := s.sessionStore.DeleteSession(ctx, sessionID); err != nil {
		return fmt.Errorf("delete session: %w", err)
	}

	return nil
}

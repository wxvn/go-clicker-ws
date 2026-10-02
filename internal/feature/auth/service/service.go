package service

import (
	"context"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"github.com/wxvn/go-clicker-ws/internal/hasher"
	"github.com/wxvn/go-clicker-ws/internal/session"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}

type AuthService struct {
	userRepository UserRepository
	hasher         hasher.PasswordHasher
	sessionStore   session.Store
}

func NewAuthService(userRepository UserRepository, passwordHasher hasher.PasswordHasher, sessionStore session.Store) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		hasher:         passwordHasher,
		sessionStore:   sessionStore,
	}
}

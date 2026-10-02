package http

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/wxvn/go-clicker-ws/internal/domain"
)

type Service interface {
	Register(ctx context.Context, auth domain.Auth) (domain.User, string, error)
	Login(ctx context.Context, auth domain.Auth) (domain.User, string, error)
	Logout(ctx context.Context, sessionID string) error
}

type Handler struct {
	service  Service
	validate *validator.Validate
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service:  service,
		validate: validator.New(),
	}
}

type AuthRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type AuthResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   int    `json:"avatar"`
	Clicks   int64  `json:"clicks"`
}

package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wxvn/go-clicker-ws/internal/domain"
)

type ClicksService interface {
	GetUser(ctx context.Context, userId string) (domain.User, error)
	GetLeaderboard(ctx context.Context, limit int, userID string) (domain.LeaderboardResult, error)
	IncrementClicks(ctx context.Context, userID string) (int64, error)
}

type ClicksHandler struct {
	service  ClicksService
	validate *validator.Validate
}

func NewHandler(service ClicksService) *ClicksHandler {
	return &ClicksHandler{
		service:  service,
		validate: validator.New(),
	}
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   int    `json:"avatar"`
	Clicks   int64  `json:"clicks"`
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}

type LeaderboardResponse struct {
	Users       []LeaderboardUserResponse `json:"users"`
	CurrentUser *LeaderboardUserResponse  `json:"current_user,omitempty"`
}

type LeaderboardUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   int    `json:"avatar"`
	Clicks   int64  `json:"clicks"`
	Position *int64 `json:"position,omitempty"`
}

type ClickResponse struct {
	Clicks int64 `json:"clicks"`
}

package ws

import (
	"context"
	"log"
	"net/http"
	"slices"

	"github.com/gorilla/websocket"
	"github.com/wxvn/go-clicker-ws/internal/domain"
	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

type Service interface {
	// GetUser(ctx context.Context, userID string) (domain.User, error)
	GetLeaderboard(ctx context.Context, limit int, userID string) (domain.LeaderboardResult, error)
	IncrementClicks(ctx context.Context, userID string) (int64, error)
}

type Handler struct {
	hub            *Hub
	service        Service
	allowedOrigins []string
}

func NewHandler(hub *Hub, service Service, allowedOrigins []string) *Handler {
	return &Handler{
		hub:            hub,
		service:        service,
		allowedOrigins: allowedOrigins,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return slices.Contains(h.allowedOrigins, origin)
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:     h.hub,
		conn:    conn,
		send:    make(chan []byte, 256),
		userID:  userID,
		service: h.service,
	}

	h.hub.register <- client

	go client.writePump()
	go client.readPump()
}

package http

import (
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/session"
	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

func (h *ClicksHandler) RegisterRoutes(router *http.ServeMux, sessionStore session.Store) {
	auth := middleware.Auth(sessionStore)

	router.Handle("GET /user/{id}", auth(http.HandlerFunc(h.GetUser)))
	router.Handle("GET /user/me", auth(http.HandlerFunc(h.GetMeUser)))
	router.Handle("GET /leaderboard", auth(http.HandlerFunc(h.GetLeaderboard)))
	router.Handle("POST /click", auth(http.HandlerFunc(h.Click)))
}

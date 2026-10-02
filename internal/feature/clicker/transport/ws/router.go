package ws

import (
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/session"
	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

func (h *Handler) RegisterRoutes(router *http.ServeMux, sessionStore session.Store) {
	router.Handle("/ws",
		middleware.Auth(sessionStore)(http.HandlerFunc(h.ServeHTTP)))
	//middleware.Auth(sessionStore)(h))
}

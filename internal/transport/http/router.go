package http

import (
	"log/slog"
	"net/http"

	authhttp "github.com/wxvn/go-clicker-ws/internal/feature/auth/transport/http"
	clickhttp "github.com/wxvn/go-clicker-ws/internal/feature/clicker/transport/http"
	"github.com/wxvn/go-clicker-ws/internal/feature/clicker/transport/ws"
	"github.com/wxvn/go-clicker-ws/internal/session"
	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func NewRouter(authHandler *authhttp.Handler, clickHandler *clickhttp.ClicksHandler, wsHandler *ws.Handler, log *slog.Logger, sessionStore session.Store, allowedOrigins []string) http.Handler {
	mux := http.NewServeMux()

	authHandler.RegisterRoutes(mux)
	clickHandler.RegisterRoutes(mux, sessionStore)
	wsHandler.RegisterRoutes(mux, sessionStore)

	wrappedMux := Chain(mux, middleware.CORS(allowedOrigins), middleware.Recovery(log), middleware.RequestId, middleware.Logger(log))

	return wrappedMux

}

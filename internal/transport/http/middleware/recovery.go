package middleware

import (
	"log/slog"
	"net/http"
)

func Recovery(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					log.Error(
						"panic recovered",
						"method", r.Method,
						"path", r.URL.Path,
						"request_id", GetRequestID(r.Context()),
						"panic", p,
					)

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write([]byte(`{"error":"something went wrong"}`))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

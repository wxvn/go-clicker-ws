package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/session"
)

type userIDKey struct{}

func withUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func GetUserID(ctx context.Context) string {
	userID, _ := ctx.Value(userIDKey{}).(string)
	return userID
}

func Auth(sessionStore session.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("AUTH MIDDLEWARE:", r.Method, r.URL.Path)
			cookie, err := r.Cookie("session_id")
			if err != nil {
				writeUnauthorized(w)
				return
			}
			log.Println("COOKIE:", cookie.Value)

			userID, err := sessionStore.GetSession(
				r.Context(),
				cookie.Value,
			)
			if err != nil {
				writeUnauthorized(w)
				return
			}

			log.Println("USER ID:", userID)

			err = sessionStore.RefreshSession(
				r.Context(),
				cookie.Value,
			)
			if err != nil {
				writeUnauthorized(w)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    cookie.Value,
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				MaxAge:   60 * 60 * 24 * 7,
				SameSite: http.SameSiteLaxMode,
			})

			r = r.WithContext(
				withUserID(r.Context(), userID),
			)

			next.ServeHTTP(w, r)
		})
	}
}

func writeUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	data := map[string]string{
		"error": "unauthorized",
	}

	_ = json.NewEncoder(w).Encode(data)
}

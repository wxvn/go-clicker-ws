package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const requestIdKey = "requestId"

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()

		ctx := r.Context()

		ctx = context.WithValue(ctx, requestIdKey, id)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIdKey).(string); ok {
		return id
	}
	return ""
}

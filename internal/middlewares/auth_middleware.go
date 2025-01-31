package middlewares

import (
	"context"
	"net/http"
)

type key string

const (
	userDataKey key = "userData"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData := "some user data"

		ctx := context.WithValue(r.Context(), userDataKey, userData)

		// Pass the new context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserData(ctx context.Context) string {
	res, ok := ctx.Value(userDataKey).(string)
	if !ok {
		return ""
	}
	return res
}

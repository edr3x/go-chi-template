package middlewares

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func ZapLoggerMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(rw, r)
			duration := time.Since(start)

			if rw.statusCode < 400 {
				zap.L().WithOptions(zap.WithCaller(false)).Info("Request Completed",
					zap.String("Method", r.Method),
					zap.String("Path", r.URL.Path),
					zap.String("Remote", r.RemoteAddr),
					zap.Int("Status", rw.statusCode),
					zap.Duration("Duration", duration),
				)
			}
		}
		return http.HandlerFunc(fn)
	}
}

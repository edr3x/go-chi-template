package responders

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/edr3x/chi-template/internal/entities"
)

type failureResponse struct {
	Success bool `json:"success"`
	Message any  `json:"message"`
}

func (s *responders) jsonError(w http.ResponseWriter, e any, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(failureResponse{
		Success: false,
		Message: e,
	})
}

func (s *responders) Error(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case entities.HttpError:
		s.logger.Error(
			err.Error(),
			zap.Int("Status", e.Code),
			zap.String("Method", r.Method),
			zap.String("Caller", e.Caller),
			zap.String("URI", r.URL.RequestURI()),
			zap.String("Request-ID", middleware.GetReqID(r.Context())),
		)

		if e.Code >= 500 {
			s.jsonError(w, "Internal Server Error", e.Code)
			return
		}
		s.jsonError(w, err.Error(), e.Code)
		return
	default:
		s.logger.Error(
			err.Error(),
			zap.String("Method", r.Method),
			zap.String("URI", r.URL.RequestURI()),
			zap.String("Request-ID", middleware.GetReqID(r.Context())),
		)
		s.jsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

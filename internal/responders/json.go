package responders

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success bool `json:"success"`
	Payload any  `json:"payload"`
}

func (s *responders) JSON(w http.ResponseWriter, i any, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(successResponse{
		Success: true,
		Payload: i,
	})
}

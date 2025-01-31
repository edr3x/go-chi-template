package v1

import (
	"github.com/go-chi/chi/v5"

	"github.com/edr3x/chi-template/internal/handlers"
)

func AuthRoutes(r chi.Router) {
	handler := handlers.NewAuthHandler()
	r.Post("/login", handler.Login)
}

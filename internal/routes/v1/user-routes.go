package v1

import (
	"github.com/edr3x/chi-template/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	handler := handlers.NewUserHandler()

	r.Get("/{id}", handler.GetUsers)
}

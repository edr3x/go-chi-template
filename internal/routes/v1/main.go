package v1

import (
	"github.com/go-chi/chi/v5"

	"github.com/edr3x/chi-template/internal/middlewares"
)

func MainRouter(r chi.Router) {
	r.Route("/auth", AuthRoutes)
	r.Route("/users", func(rr chi.Router) {
		rr.Use(middlewares.RequireAuth)
		UserRoutes(rr)
	})
}

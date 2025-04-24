package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/edr3x/chi-template/internal/entities"
	"github.com/edr3x/chi-template/internal/middlewares"
	"github.com/edr3x/chi-template/internal/responders"
	v1 "github.com/edr3x/chi-template/internal/routes/v1"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) NewHandler() http.Handler {
	app := chi.NewRouter()

	res := responders.NewResponders()

	// Middlewares
	app.Use(middlewares.ZapLoggerMiddleware())
	app.Use(middleware.CleanPath)
	app.Use(middleware.RequestID)

	app.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	// Security Headers
	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Frame-Options", "SAMEORIGIN")
			w.Header().Set("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
			next.ServeHTTP(w, r)
		})
	})

	app.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		res.JSON(w, "Hello there")
	})

	// Routes
	app.Route("/api/v1", v1.MainRouter)

	// 404 Handler
	app.NotFound(func(w http.ResponseWriter, r *http.Request) {
		res.Error(w, r, entities.ErrorNotFound("Route Not Found"))
	})

	// Method Not Allowed
	app.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		res.Error(w, r, entities.ErrorMethodNotAllowed())
	})

	return app
}

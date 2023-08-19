package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	router "github.com/edr3x/chi-explore/modules"
	"github.com/edr3x/chi-explore/utils"
)

type FailureResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	app := chi.NewRouter()
	app.Use(middleware.Heartbeat("/"))

	app.Use(middleware.Logger)

	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					statusCode := http.StatusInternalServerError
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("internal server error")
					}
					if e, ok := err.(*utils.ErrorTypeStruct); ok {
						statusCode = e.StatusCode
					}
					failureResponse := FailureResponse{
						Success: false,
						Message: err.Error(),
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(statusCode)
					json.NewEncoder(w).Encode(failureResponse)
				}
			}()
			next.ServeHTTP(w, r)
		})
	})

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

	// Routes
	app.Route("/api/v1", router.MainRouter)

	// 404 Handler
	app.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		panic(utils.CustomError(404, "Endpoint Not Found"))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	log.Println("listening in port: " + port + "...")
	http.ListenAndServe("0.0.0.0:"+port, app)
}

package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	// list of trusted proxies
	TrustedProxies []string

	// The address the server listens on
	ListenAddress string
}

// Returns a configured gin router with all routes registered
func NewServer(config Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return router
}

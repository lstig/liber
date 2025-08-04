package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"

	"github.com/lstig/liber/internal/handlers"
	"github.com/lstig/liber/web"
)

type Option func(s *Server) error

func WithLogger(logger *httplog.Logger) Option {
	return func(s *Server) error {
		s.logger = logger
		return nil
	}
}

func WithMiddleware(middlewares ...func(http.Handler) http.Handler) Option {
	return func(s *Server) error {
		s.router.Use(middlewares...)
		return nil
	}
}

type Server struct {
	logger *httplog.Logger
	router *chi.Mux
}

func NewServer(opts ...Option) (*Server, error) {
	// configure the default server
	s := &Server{}
	s.router = chi.NewRouter()

	// apply options
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	// add a default logger if one wasn't provided
	if s.logger == nil {
		s.logger = httplog.NewLogger("server")
	}

	return s, nil
}

func (s *Server) mountHandlers() {
	// instantiate services
	health := handlers.NewHealthHandler(s.logger)
	home := handlers.NewHomeHandler(s.logger)

	// endpoints
	s.router.Get("/", home.ServeHTTP)

	// health check
	s.router.Get("/health", health.ServeHTTP)

	// static assets
	s.router.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r) })
	s.router.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r) })

}

func (s *Server) ListenAndServe(addr string) error {
	s.logger.Info("server starting", "log_level", s.logger.Options.LogLevel)
	s.mountHandlers()
	s.logger.Info("server listening", "address", addr)
	return http.ListenAndServe(addr, s.router)
}

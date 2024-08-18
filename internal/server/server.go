package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"

	"github.com/lstig/liber/internal/handlers"
	"github.com/lstig/liber/internal/middleware"
	"github.com/lstig/liber/internal/views"
	"github.com/lstig/liber/web"
)

type Option func(s *Server) error

func WithLogger(logger *httplog.Logger) Option {
	return func(s *Server) error {
		s.logger = logger
		return nil
	}
}

func WithDevMode(enabled bool) Option {
	return func(s *Server) error {
		s.devMode = enabled
		return nil
	}
}

func WithProperties(props *views.Properties) Option {
	return func(s *Server) error {
		s.viewProps = props
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
	logger    *httplog.Logger
	router    *chi.Mux
	devMode   bool
	viewProps *views.Properties
}

func NewServer(opts ...Option) (*Server, error) {
	// configure the default server
	s := &Server{
		viewProps: &views.Properties{},
	}
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

	// endpoints
	s.router.Get("/", templ.Handler(views.Home(s.viewProps)).ServeHTTP)

	// health check
	s.router.Get("/health", health.Health)

	// static assets
	s.router.Group(func(r chi.Router) {
		if s.devMode {
			r.Use(middleware.SetHeader("Cache-Control", "no-cache, no-store, must-revalidate"))
		}
		r.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r) })
		r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r) })
	})
}

func (s *Server) ListenAndServe(addr string) error {
	s.logger.Info("server starting", "log_level", s.logger.Options.LogLevel, "dev_mode", s.devMode)
	s.mountHandlers()
	s.logger.Info("server listening", "address", addr)
	return http.ListenAndServe(addr, s.router)
}

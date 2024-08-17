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
		s.Logger = logger
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
		s.Router.Use(middlewares...)
		return nil
	}
}

type Server struct {
	Router    *chi.Mux
	Logger    *httplog.Logger
	devMode   bool
	viewProps *views.Properties
}

func NewServer(opts ...Option) (*Server, error) {
	// configure the default server
	s := &Server{
		viewProps: &views.Properties{},
	}
	s.Router = chi.NewRouter()

	// apply options
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	// add a default logger if one wasn't provided
	if s.Logger == nil {
		s.Logger = httplog.NewLogger("server")
	}

	return s, nil
}

func (s *Server) MountHandlers() {
	// instantiate services
	health := handlers.NewHealthHandler(s.Logger)

	// endpoints
	s.Router.Get("/", templ.Handler(views.Home(s.viewProps)).ServeHTTP)

	// health check
	s.Router.Get("/health", health.Health)

	// static assets
	s.Router.Group(func(r chi.Router) {
		if s.devMode {
			r.Use(middleware.SetHeader("Cache-Control", "no-cache, no-store, must-revalidate"))
		}
		r.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r) })
		r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r) })
	})
}

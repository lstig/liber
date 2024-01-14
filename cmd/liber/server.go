package main

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/lstig/liber/handlers"
	"github.com/lstig/liber/middleware"
	"github.com/lstig/liber/views"
	"github.com/lstig/liber/web"
	"github.com/spf13/cobra"
)

type Server struct {
	ListenAddress string
	Dev           bool
	Router        *chi.Mux
	Logger        *httplog.Logger
	viewProps     *views.Properties
}

func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Logger = httplog.NewLogger("liber", httplog.Options{
		Concise:         true,
		TimeFieldFormat: time.RFC3339,
	})
	s.viewProps = &views.Properties{}
	return s
}

// BindFlags register flags and bind them to the fields of the 'server' struct
func (s *Server) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&s.ListenAddress, "listen", "l", ":8080", "the server's listening address")
	cmd.Flags().BoolVar(&s.Dev, "dev", false, "run server with additional configuration for development")
}

func (s *Server) configureMiddleware() {
	s.Logger.Debug("configuring middleware")

	// dev mode only middlware
	if s.Dev {
		s.Router.Use(middleware.Prefer)
	}

	s.Router.Use(chimiddleware.RequestID)
	s.Router.Use(httplog.Handler(s.Logger, []string{"/health"}))
	s.Router.Use(chimiddleware.Recoverer)
}

func (s *Server) mountHandlers() {
	s.Logger.Debug("initializing services")
	health := handlers.NewHealthHandler(s.Logger)

	s.Logger.Debug("registering routes")
	s.Router.Get("/", templ.Handler(views.Home(s.viewProps)).ServeHTTP)
	s.Router.Get("/health", health.Health)
	s.Router.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r)
	})
	s.Router.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r)
	})
}

func (s *Server) setProperties() {
	s.viewProps.Dev = s.Dev
}

// Run configures the router and starts the server on the specified address/port
func (s *Server) Run(_ *cobra.Command, _ []string) error {
	s.Logger.Info("server starting", "log_level", s.Logger.Options.LogLevel, "dev_mode", s.Dev)
	s.setProperties()
	s.configureMiddleware()
	s.mountHandlers()
	s.Logger.Info("server listening", "address", s.ListenAddress)
	return http.ListenAndServe(s.ListenAddress, s.Router)
}

// newServerCommand returns a configured server command
func newServerCommand() *cobra.Command {
	s := NewServer()
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run eBook server",
		RunE:  s.Run,
	}
	s.BindFlags(cmd)
	return cmd
}

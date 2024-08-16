package main

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/lstig/liber/internal/handlers"
	"github.com/lstig/liber/internal/middleware"
	"github.com/lstig/liber/internal/views"
	"github.com/lstig/liber/web"
	"github.com/spf13/cobra"
)

type ServerOption func(s *Server) error

func WithLogger(logger *httplog.Logger) ServerOption {
	return func(s *Server) error {
		s.Logger = logger
		return nil
	}
}

func WithProperties(props *views.Properties) ServerOption {
	return func(s *Server) error {
		s.viewProps = props
		return nil
	}
}

func WithMiddlware(middlewares ...func(http.Handler) http.Handler) ServerOption {
	return func(s *Server) error {
		s.Router.Use(middlewares...)
		return nil
	}
}

type Server struct {
	Router    *chi.Mux
	Logger    *httplog.Logger
	viewProps *views.Properties
}

func NewServer(opts ...ServerOption) (*Server, error) {
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

func (s *Server) mountHandlers() {
	// instantiate services
	health := handlers.NewHealthHandler(s.Logger)

	// endpoints
	s.Router.Get("/", templ.Handler(views.Home(s.viewProps)).ServeHTTP)

	// health check
	s.Router.Get("/health", health.Health)

	// static assets
	s.Router.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r)
	})
	s.Router.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r)
	})
}

func (cli *CLI) serverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run eBook server",
		RunE:  cli.serverRun,
	}

	cmd.Flags().StringVarP(&cli.server.listenAddress, "listen", "l", ":8081", "the server's listening address")

	return cmd
}

func (cli *CLI) serverRun(cmd *cobra.Command, args []string) error {
	logger := httplog.NewLogger("liber", httplog.Options{
		Concise:         true,
		LogLevel:        httplog.LevelByName(cli.verbosity),
		TimeFieldFormat: time.RFC3339,
	})
	srv, err := NewServer(
		WithLogger(logger),
		WithProperties(&views.Properties{
			Dev: cli.devMode,
		}),
		WithMiddlware(
			chimiddleware.RequestID,
			httplog.Handler(logger, []string{"/health"}),
			chimiddleware.Recoverer,
		),
	)
	if err != nil {
		return err
	}

	// add dev specific middleware
	if cli.devMode {
		// adding to the front of the chain so we're called before Recoverer
		middlewares := append(chi.Middlewares{middleware.Prefer}, srv.Router.Middlewares()...)
		srv.Router.Use(middlewares...)
	}

	srv.Logger.Info("server starting", "log_level", srv.Logger.Options.LogLevel, "dev_mode", cli.devMode)
	srv.mountHandlers()
	srv.Logger.Info("server listening", "address", cli.server.listenAddress)
	return http.ListenAndServe(cli.server.listenAddress, srv.Router)
}

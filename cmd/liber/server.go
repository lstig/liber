package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lstig/liber/views"
	"github.com/lstig/liber/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Server struct {
	ListenAddress string
	Router        *chi.Mux
	Logger        *logrus.Logger
}

func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Logger = logrus.New()
	return s
}

// BindFlags register flags and bind them to the fields of the 'server' struct
func (s *Server) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&s.ListenAddress, "listen", "l", ":8080", "the server's listening address")
}

func (s *Server) MountHandlers() {
	// add middlewares
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// add handlers
	s.Router.Get("/", templ.Handler(views.Home()).ServeHTTP)
	s.Router.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r)
	})
	s.Router.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r)
	})
}

// Run configures the router and starts the server on the specified address/port
func (s *Server) Run(_ *cobra.Command, _ []string) error {
	s.Logger.Infof("server listening on address %s", s.ListenAddress)
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

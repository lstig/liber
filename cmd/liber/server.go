package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lstig/liber/views"
	"github.com/lstig/liber/web"
	"github.com/spf13/cobra"
)

type server struct {
	ListenAddress string
}

// BindFlags register flags and bind them to the fields of the 'server' struct
func (s *server) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&s.ListenAddress, "listen", "l", ":8080", "the server's listening address")
}

// Run configures the router and starts the server on the specified address/port
func (s *server) Run(_ *cobra.Command, _ []string) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", templ.Handler(views.Home()).ServeHTTP)
	r.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r)
	})
	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r)
	})

	slog.Info("server listening on " + s.ListenAddress)
	return http.ListenAndServe(s.ListenAddress, r)
}

// newServerCommand returns a configured server command
func newServerCommand() *cobra.Command {
	srv := &server{}
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run eBook server",
		RunE:  srv.Run,
	}
	srv.BindFlags(cmd)
	return cmd
}

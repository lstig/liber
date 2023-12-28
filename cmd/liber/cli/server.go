package cli

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var listenAddress string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run eBook server",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := chi.NewRouter()

		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("."))
		})

		slog.Info("server listening", "address", fmt.Sprintf("http://%v", listenAddress))
		return http.ListenAndServe(listenAddress, r)
	},
}

func init() {
	serverCmd.Flags().StringVarP(&listenAddress, "address", "a", "127.0.0.1:8080", "the server's listening address")

	rootCmd.AddCommand(serverCmd)
}

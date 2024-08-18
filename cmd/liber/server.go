package main

import (
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/spf13/cobra"

	"github.com/lstig/liber/internal/server"
	"github.com/lstig/liber/internal/views"
)

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
	srv, err := server.NewServer(
		server.WithLogger(logger),
		server.WithDevMode(cli.devMode),
		server.WithProperties(&views.Properties{
			Dev: cli.devMode,
		}),
		server.WithMiddleware(
			chimiddleware.RequestID,
			httplog.Handler(logger, []string{"/health"}),
			chimiddleware.Recoverer,
		),
	)
	if err != nil {
		return err
	}

	return srv.ListenAndServe(cli.server.listenAddress)
}

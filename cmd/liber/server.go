package main

import (
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/spf13/cobra"

	"github.com/lstig/liber/internal/server"
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

func (cli *CLI) serverRun(_ *cobra.Command, _ []string) error {
	logger := httplog.NewLogger("liber", httplog.Options{
		Concise:         true,
		LogLevel:        httplog.LevelByName(cli.verbosity),
		TimeFieldFormat: time.RFC3339,
	})
	srv, err := server.NewServer(
		server.WithLogger(logger),
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

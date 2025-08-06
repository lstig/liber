package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/go-chi/httplog/v2"
	"github.com/spf13/cobra"

	"github.com/lstig/liber/internal/router"
)

type ServerOptions struct {
	ListenAddr      string
	LogLevel        string
	ShutdownTimeout time.Duration
}

func serverCmd(ctx context.Context) *cobra.Command {
	opts := &ServerOptions{}
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run eBook server",
		RunE:  func(_ *cobra.Command, _ []string) error { return runServer(ctx, opts) },
	}

	cmd.Flags().StringVarP(&opts.ListenAddr, "listen", "l", "localhost:8081", "the server's listening address")
	cmd.Flags().StringVar(&opts.LogLevel, "log-level", slog.LevelInfo.String(), "log level")
	cmd.Flags().DurationVar(&opts.ShutdownTimeout, "shutdown-timeout", 15*time.Second, "shutdown timeout")

	return cmd
}

func runServer(ctx context.Context, opts *ServerOptions) error {
	var (
		shuttingDown atomic.Bool
		logger       = httplog.NewLogger("liber", httplog.Options{
			Concise:         true,
			LogLevel:        httplog.LevelByName(opts.LogLevel),
			TimeFieldFormat: time.RFC3339,
		})
	)

	// router
	r, err := router.NewRouter(
		logger,
		router.WithMiddleware(httplog.RequestLogger(logger, []string{"/healthz"})),
		router.WithRoute("GET", "/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if shuttingDown.Load() {
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
				return
			}
			_, err := fmt.Fprintln(w, "ok")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		})),
	)
	if err != nil {
		return err
	}

	// server
	requestCtx, drainRequests := context.WithCancel(ctx)
	srv := &http.Server{
		Addr:    opts.ListenAddr,
		Handler: r,
		BaseContext: func(_ net.Listener) context.Context {
			return requestCtx
		},
	}
	go func() {
		logger.Info("starting server", "addr", opts.ListenAddr, "log_level", opts.LogLevel)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	// shutdown
	<-ctx.Done()
	shuttingDown.Store(true)
	logger.Info("received shutdown signal, shutting down")
	logger.Info("waiting for requests to finish", "timeout", opts.ShutdownTimeout.String())
	shutdownCtx, shutdown := context.WithTimeout(ctx, opts.ShutdownTimeout)
	defer shutdown()
	err = srv.Shutdown(shutdownCtx)
	drainRequests()
	if err != nil {
		return err
	}

	logger.Info("server shutdown successfully!")

	return nil
}

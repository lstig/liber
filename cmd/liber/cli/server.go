package cli

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/lstig/liber/internal/server"
)

var config server.Config

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run eBook server",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := server.NewServer(config)
		slog.Info("server listening", "address", fmt.Sprintf("http://%v", config.ListenAddress))
		err := http.ListenAndServe(config.ListenAddress, srv)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	serverCmd.Flags().StringVarP(&config.ListenAddress, "address", "a", "127.0.0.1:8080", "the server's listening address")
	serverCmd.Flags().StringSliceVar(&config.TrustedProxies, "trusted-proxies", nil, "comma separated list of IPs that are trusted proxies")

	rootCmd.AddCommand(serverCmd)
}

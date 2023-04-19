package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lstig/liber/internal/server"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run eBook server",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := server.NewServer()
		err := srv.Run()
		if err != nil {
			return err
		}
		return nil
	},
}

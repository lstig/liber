package main

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func (cli *CLI) root() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "liber",
		Short:             "Liber is an eBook server with OPDS support",
		Version:           Version,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	cmd.PersistentFlags().StringVarP(&cli.verbosity, "verbosity", "v", slog.LevelInfo.String(), "log level to display (DEBUG, INFO, WARN, or ERROR)")
	cmd.PersistentFlags().BoolVar(&cli.devMode, "dev", false, "run cli with additional configuration for development")

	cmd.AddCommand(
		cli.serverCmd(),
	)

	return cmd
}

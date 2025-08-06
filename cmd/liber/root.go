package main

import (
	"context"

	"github.com/spf13/cobra"
)

func rootCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "liber",
		Short:             "Liber is an eBook server with OPDS support",
		Version:           Version,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	cmd.AddCommand(
		serverCmd(ctx),
	)

	return cmd
}

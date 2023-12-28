package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string = "not built correctly"

var rootCmd = &cobra.Command{
	Use:               "liber",
	Short:             "Liber is an eBook server with OPDS support",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func main() {
	rootCmd.AddCommand(newServerCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

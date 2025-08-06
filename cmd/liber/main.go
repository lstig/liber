package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var Version string = "not built correctly"

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	return rootCmd(ctx).Execute()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
)

var Version string = "not built correctly"

type CLI struct {
	verbosity string
	server    struct {
		listenAddress string
	}
}

func main() {
	cli := &CLI{}

	if err := cli.root().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package server

import (
	// "github.com/gofrs/uuid"
	"fmt"
	"net/http"
)

var Version string = "not built correctly"

type Config struct {
	// list of trusted proxies
	TrustedProxies []string

	// The address the server listens on
	ListenAddress string
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Index\n")
}

// Returns a configured gin router with all routes registered
func NewServer(config Config) *http.ServeMux {
	http.HandleFunc("/", index)

	return http.DefaultServeMux
}

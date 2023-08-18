package server

import (
	// "github.com/gofrs/uuid"
	"os"
	"log/slog"
	"net/http"

	"github.com/lstig/liber/internal/middleware"
	"github.com/lstig/liber/web"
)

var Version string = "not built correctly"

type Config struct {
	// list of trusted proxies
	TrustedProxies []string

	// The address the server listens on
	ListenAddress string
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Index\n"))
}

// Returns a configured router with all routes registered
func NewServer(config Config) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.Handle("/assets/", http.FileServer(http.FS(web.Assets)))

	logging := middleware.Logging(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	router := logging(mux)

	return router
}


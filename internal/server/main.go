package server

import (
	// "github.com/gofrs/uuid"
	"html/template"
	"log/slog"
	"net/http"
	"os"

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

type SiteConfig struct {
	AssetPath string
	Title     string
}

func index(w http.ResponseWriter, req *http.Request) {
	cfg := &SiteConfig{AssetPath: "assets", Title: "Liber"}
	view := template.Must(template.ParseFS(web.Views, "views/base.gohtml", "views/index.gohtml"))
	view.ExecuteTemplate(w, "index.gohtml", cfg)
}

// NewServer returns a configured server with all routes registered
func NewServer(config Config) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/assets/", http.FileServer(http.FS(web.Assets)))

	logging := middleware.Logging(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	router := logging(mux)

	return router
}

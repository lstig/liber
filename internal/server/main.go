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
	Title     string
	Page      string
}

var config SiteConfig = SiteConfig{
	Title: "Liber",
}

func index(w http.ResponseWriter, req *http.Request) {
	config.Page = "Home"
	view := template.Must(template.ParseFS(web.FS, "views/*.gohtml"))
	view.ExecuteTemplate(w, "index.gohtml", config)
}

func books(w http.ResponseWriter, req *http.Request) {
	config.Page = "Books"
	view := template.Must(template.ParseFS(web.FS, "views/*.gohtml"))
	view.ExecuteTemplate(w, "index.gohtml", config)
}

func collections(w http.ResponseWriter, req *http.Request) {
	config.Page = "Collections"
	view := template.Must(template.ParseFS(web.FS, "views/*.gohtml"))
	view.ExecuteTemplate(w, "index.gohtml", config)
}

// NewServer returns a configured server with all routes registered
func NewServer(config Config) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/books", books)
	mux.HandleFunc("/collections", collections)
	mux.Handle("/assets/", http.FileServer(http.FS(web.FS)))

	logging := middleware.Logging(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	router := logging(mux)

	return router
}

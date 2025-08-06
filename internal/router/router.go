package router

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"

	"github.com/lstig/liber/internal/views"
	"github.com/lstig/liber/web"
)

type Option func(r *chi.Mux)

func WithMiddleware(middleware ...func(http.Handler) http.Handler) Option {
	return func(r *chi.Mux) {
		r.Use(middleware...)
	}
}

func WithRoute(method, pattern string, handler http.Handler) Option {
	return func(r *chi.Mux) {
		r.Method(method, pattern, handler)
	}
}

func NewRouter(logger *httplog.Logger, options ...Option) (*chi.Mux, error) {
	r := chi.NewRouter()

	for _, opt := range options {
		opt(r)
	}

	// core routes
	r.Get("/", servePage(logger, views.Home()))
	r.Get("/books", servePage(logger, views.Books()))

	// static assets
	r.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Dist)).ServeHTTP(w, r) })
	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) { http.FileServer(http.FS(web.Assets)).ServeHTTP(w, r) })

	return r, nil
}

func servePage(log *httplog.Logger, page templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "" {
			if err := page.Render(r.Context(), w); err != nil {
				log.Error(fmt.Sprintf("error during response: %s", err))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		if err := templ.RenderFragments(r.Context(), w, page, "page"); err != nil {
			log.Error(fmt.Sprintf("error during response: %s", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

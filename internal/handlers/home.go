package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/httplog/v2"

	"github.com/lstig/liber/internal/views"
)

type HomeHandler struct {
	Log *httplog.Logger
}

func NewHomeHandler(log *httplog.Logger) *HomeHandler {
	return &HomeHandler{log}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *HomeHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.View(w, r)
}

func (h *HomeHandler) View(w http.ResponseWriter, r *http.Request) {
	if err := views.Home().Render(r.Context(), w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.Error(fmt.Sprintf("error during response: %s", err))
	}
}

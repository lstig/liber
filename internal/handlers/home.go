package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/httplog/v2"

	"github.com/lstig/liber/internal/views"
)

type HomeHandler struct {
	Log   *httplog.Logger
	props *views.GlobalProperties
}

func NewHomeHandler(log *httplog.Logger, props *views.GlobalProperties) *HomeHandler {
	return &HomeHandler{log, props}
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
	if err := views.Home(h.props).Render(r.Context(), w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.Error(fmt.Sprintf("error during response: %s", err))
	}
}

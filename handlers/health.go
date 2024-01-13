package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

type HealhtHandler struct {
	Log *httplog.Logger
}

func NewHealthHandler(log *httplog.Logger) *HealhtHandler {
	return &HealhtHandler{
		Log: log,
	}
}

func (h *HealhtHandler) Health(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("ok")); err != nil {
		h.Log.Error(fmt.Sprintf("could not write data to response: %s", err))
	}
}

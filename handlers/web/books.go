package book

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/httplog/v2"
	"github.com/lstig/liber/services/book"
	"github.com/lstig/liber/views"
)

type BookService interface {
	Get(ctx context.Context, id string) (*book.Book, error)
	List(ctx context.Context) ([]book.Book, error)
}

type BookHandler struct {
	Log *httplog.Logger
	BookService
}

func NewBookHandler(log *httplog.Logger, bs BookService) *BookHandler {
	return &BookHandler{
		Log:         log,
		BookService: bs,
	}
}

func (h *BookHandler) Books(w http.ResponseWriter, r *http.Request) {
	books, err := h.List(r.Context())
	if err != nil {
		// log the actual error
		h.Log.Error(fmt.Sprintf("could not retrieve books: %s", err))
		// write the response
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("could not retrieve books")); err != nil {
			h.Log.Error(fmt.Sprintf("could not write data to response: %s", err))
		}
		return
	}
	if err := views.Books(books).Render(r.Context(), w); err != nil {
		h.Log.Error(fmt.Sprintf("could not write data to response: %s", err))
	}
}

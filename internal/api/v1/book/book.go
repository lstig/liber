package book

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/usecase"
)

type Book struct {
	Title      string `json:"title"`
	Identifier string `json:"identifier"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
}

func RegisterRoutes(router chi.Router) {
	router.Method(http.MethodPost, "/book", nethttp.NewHandler(create()))
}

type createOutput struct {
	ID uuid.UUID `json:"id" type:"uuid"`
}

func create() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input Book, output *createOutput) error {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		output.ID = id
		return nil
	})
	u.SetTags("Book")

	return u
}

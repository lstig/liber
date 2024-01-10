package book

import (
	"context"

	"github.com/go-chi/httplog/v2"
)

type Book struct {
	ID     string
	Title  string
	Author string
}

type BookService struct {
	Log *httplog.Logger
}

func NewBookService(log *httplog.Logger) *BookService {
	return &BookService{Log: log}
}

func (bs *BookService) Get(ctx context.Context, id string) (*Book, error) {
	return nil, nil
}

func (bs *BookService) List(ctx context.Context) ([]Book, error) {
	return nil, nil
}

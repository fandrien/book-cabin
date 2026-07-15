package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/fandrien/book-cabin/handler"
)

func NewRouter(
	searchHandler *handler.SearchHandler,
) *chi.Mux {

	r := chi.NewRouter()

	r.Post(
		"/search",
		searchHandler.Search,
	)

	return r
}

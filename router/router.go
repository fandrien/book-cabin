package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/fandrien/book-cabin/handler"
	"github.com/fandrien/book-cabin/router/middleware"
)

func NewRouter(
	searchHandler *handler.SearchHandler,
) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RateLimit)

	r.Post(
		"/search",
		searchHandler.Search,
	)

	return r
}

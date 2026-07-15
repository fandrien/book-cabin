package provider

import (
	"context"

	"github.com/fandrien/book-cabin/model"
)

type Provider interface {
	Name() string

	Search(
		ctx context.Context,
		req model.SearchRequest,
	) ([]model.Flight, error)
}

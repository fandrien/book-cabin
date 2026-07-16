package provider

import (
	"context"

	"github.com/fandrien/book-cabin/model"
)

type Provider interface {
	Name() string
	Search(ctx context.Context) ([]model.Flight, error)
}

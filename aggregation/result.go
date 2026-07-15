package aggregation

import "github.com/fandrien/book-cabin/model"

type Result struct {
	Flights   []model.Flight
	Providers []model.ProviderResult
}

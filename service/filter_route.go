package service

import (
	"strings"

	"github.com/fandrien/book-cabin/model"
)

func matchRoute(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.Origin != "" &&
		!strings.EqualFold(req.Origin, flight.Departure.Airport) {
		return false
	}

	if req.Destination != "" &&
		!strings.EqualFold(req.Destination, flight.Arrival.Airport) {
		return false
	}

	return true
}

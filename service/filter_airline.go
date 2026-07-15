package service

import (
	"strings"

	"github.com/fandrien/book-cabin/model"
)

func matchAirline(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.Airline == "" {
		return true
	}

	if strings.EqualFold(req.Airline, flight.Airline.Name) {
		return true
	}

	return false
}

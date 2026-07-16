package service

import (
	"strings"

	"github.com/fandrien/book-cabin/model"
)

func matchAirlines(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if len(req.Airlines) > 0 {
		for i := range req.Airlines {
			if strings.EqualFold(req.Airlines[i], flight.Airline.Name) {
				return true
			}
		}
		return false
	}

	return true
}

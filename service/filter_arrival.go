package service

import (
	"time"

	"github.com/fandrien/book-cabin/model"
)

func matchArrivalDate(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.ArrivalDate == "" {
		return true
	}

	t, err := time.Parse("2006-01-02", req.ArrivalDate)
	if err != nil {
		return false
	}

	arrival := time.Unix(
		flight.Departure.Timestamp,
		0,
	)

	return arrival.Year() == t.Year() &&
		arrival.Month() == t.Month() &&
		arrival.Day() == t.Day()
}

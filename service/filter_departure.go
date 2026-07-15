package service

import (
	"time"

	"github.com/fandrien/book-cabin/model"
)

func matchDepartureDate(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.DepartureDate == "" {
		return true
	}

	t, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		return false
	}

	departure := time.Unix(
		flight.Departure.Timestamp,
		0,
	)

	return departure.Year() == t.Year() &&
		departure.Month() == t.Month() &&
		departure.Day() == t.Day()
}

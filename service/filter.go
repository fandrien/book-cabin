package service

import "github.com/fandrien/book-cabin/model"

func filterFlights(
	req model.SearchRequest,
	flights []model.Flight,
) []model.Flight {

	result := make([]model.Flight, 0)

	for _, flight := range flights {

		//filter origin and destination route
		if !matchRoute(req, flight) {
			continue
		}

		if !rangePrice(req, flight) {
			continue
		}

		if !matchStops(req, flight) {
			continue
		}

		if !matchDepartureDate(req, flight) {
			continue
		}

		if !matchArrivalDate(req, flight) {
			continue
		}

		if !rangeDuration(req, flight) {
			continue
		}

		if !matchAirlines(req, flight) {
			continue
		}

		result = append(result, flight)
	}

	return result
}

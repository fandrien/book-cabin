package service

import "github.com/fandrien/book-cabin/model"

func matchStops(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.Stops != nil && flight.Stops != *req.Stops {
		return false
	}

	return true
}

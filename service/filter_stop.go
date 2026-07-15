package service

import "github.com/fandrien/book-cabin/model"

func matchStops(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.Stop != nil && flight.Stops != *req.Stop {
		return false
	}

	return true
}

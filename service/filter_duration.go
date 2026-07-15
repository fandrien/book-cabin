package service

import "github.com/fandrien/book-cabin/model"

func matchDuration(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.Stop != nil && flight.Duration.TotalMinutes != *req.Duration {
		return false
	}

	return true
}

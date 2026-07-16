package service

import (
	"github.com/fandrien/book-cabin/model"
)

func rangeDuration(
	req model.SearchRequest,
	flight model.Flight,
) bool {
	if req.MinDuration != nil &&
		flight.Duration.TotalMinutes < *req.MinDuration {
		return false
	}

	if req.MaxDuration != nil &&
		flight.Duration.TotalMinutes > *req.MaxDuration {
		return false
	}

	return true
}

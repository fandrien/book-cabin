package service

import "github.com/fandrien/book-cabin/model"

func rangePrice(
	req model.SearchRequest,
	flight model.Flight,
) bool {

	if req.MinPrice != nil &&
		flight.Price.Amount < *req.MinPrice {
		return false
	}

	if req.MaxPrice != nil &&
		flight.Price.Amount > *req.MaxPrice {
		return false
	}

	return true
}

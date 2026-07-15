package service

import "github.com/fandrien/book-cabin/model"

func matchPrice(
	req model.SearchRequest,
	f model.Flight,
) bool {

	if req.Price != nil &&
		f.Price.Amount != *req.Price {
		return false
	}

	return true
}

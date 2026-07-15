package validator

import (
	"errors"
	"strings"
	"time"

	"github.com/fandrien/book-cabin/model"
)

func ValidateSearchRequest(req model.SearchRequest) error {
	if strings.TrimSpace(req.Origin) == "" {
		return errors.New("origin is required")
	}

	if strings.TrimSpace(req.Destination) == "" {
		return errors.New("destination is required")
	}

	if strings.EqualFold(req.Origin, req.Destination) {
		return errors.New("origin and destination cannot be the same")
	}

	if strings.TrimSpace(req.DepartureDate) == "" {
		return errors.New("departure date is required")
	}

	departureDate, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		return errors.New("invalid departure date format, expected yyyy-mm-dd")

	}

	if req.ReturnDate != nil && strings.TrimSpace(*req.ReturnDate) != "" {
		returnDate, err := time.Parse("2006-01-02", *req.ReturnDate)
		if err != nil {
			return errors.New("invalid return date format, expected yyyy-mm-dd")

		}

		if !departureDate.Before(returnDate) {
			return errors.New("departure date must be before return date")

		}
	}

	if strings.TrimSpace(req.ArrivalDate) != "" {
		arrivalDate, err := time.Parse("2006-01-02", req.ArrivalDate)
		if err != nil {
			return errors.New("invalid arrival date format, expected yyyy-mm-dd")
		}

		if departureDate.After(arrivalDate) {
			return errors.New("departure date must be before return date")

		}
	}

	return nil
}

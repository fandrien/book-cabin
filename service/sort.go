package service

import (
	"sort"
	"strings"

	"github.com/fandrien/book-cabin/model"
)

const (
	SortByPrice     = "price"
	SortByDuration  = "duration"
	SortByDeparture = "departure"
	SortByArrival   = "arrival"

	SortAsc  = "asc"
	SortDesc = "desc"
)

func sortFlights(
	req model.SearchRequest,
	flights []model.Flight,
) {

	sortBy := strings.ToLower(req.SortBy)
	sortOrder := strings.ToLower(req.SortOrder)

	desc := sortOrder == SortDesc

	sort.Slice(flights, func(i, j int) bool {

		switch sortBy {

		case SortByPrice:

			if desc {
				return flights[i].Price.Amount >
					flights[j].Price.Amount
			}

			return flights[i].Price.Amount <
				flights[j].Price.Amount

		case SortByDuration:

			if desc {
				return flights[i].Duration.TotalMinutes >
					flights[j].Duration.TotalMinutes
			}

			return flights[i].Duration.TotalMinutes <
				flights[j].Duration.TotalMinutes

		case SortByDeparture:

			if desc {
				return flights[i].Departure.Timestamp >
					flights[j].Departure.Timestamp
			}

			return flights[i].Departure.Timestamp <
				flights[j].Departure.Timestamp

		case SortByArrival:

			if desc {
				return flights[i].Arrival.Timestamp >
					flights[j].Arrival.Timestamp
			}

			return flights[i].Arrival.Timestamp <
				flights[j].Arrival.Timestamp

		default:

			return bestValueScore(flights[i]) <
				bestValueScore(flights[j])
		}

	})

}

// Lower score = better flight
func bestValueScore(f model.Flight) int {

	score := 0

	// cheaper is better
	score += f.Price.Amount

	// shorter duration is better
	score += f.Duration.TotalMinutes * 1000

	// direct flight is much better
	score += f.Stops * 500000

	return score
}

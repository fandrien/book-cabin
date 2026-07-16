package comparison

import (
	"fmt"

	"github.com/fandrien/book-cabin/model"
)

func BuildComparisons(flights []model.Flight) []model.FlightComparison {

	grouped := make(map[string][]model.Flight)

	for _, f := range flights {

		key := fmt.Sprintf(
			"%s|%s|%d|%d|%s",
			f.Departure.Airport,
			f.Arrival.Airport,
			f.Departure.Timestamp,
			f.Arrival.Timestamp,
			f.Airline,
		)

		grouped[key] = append(grouped[key], f)
	}

	comparisons := make([]model.FlightComparison, 0)

	for key, flights := range grouped {

		if len(flights) < 2 {
			continue
		}

		bestPrice := flights[0].Price.Amount
		worstPrice := flights[0].Price.Amount

		offers := make([]model.Offer, 0)

		for _, f := range flights {

			price := f.Price.Amount

			if price < bestPrice {
				bestPrice = price
			}

			if price > worstPrice {
				worstPrice = price
			}

			offers = append(offers, model.Offer{
				Provider: f.Provider,
				Price:    price,
			})
		}

		comparisons = append(comparisons, model.FlightComparison{
			Key:       key,
			BestPrice: bestPrice,
			Savings:   worstPrice - bestPrice,
			Offers:    offers,
		})
	}

	return comparisons
}

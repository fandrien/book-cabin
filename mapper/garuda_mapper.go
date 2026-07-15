package mapper

import (
	"fmt"
	"time"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/model"
	"github.com/fandrien/book-cabin/util"
)

func MapGaruda(resp dto.GarudaResponse) ([]model.Flight, error) {
	flights := make([]model.Flight, 0, len(resp.Flights))

	for _, flight := range resp.Flights {
		departureTime, err := time.Parse(time.RFC3339, flight.Departure.Time)
		if err != nil {
			return nil, err
		}
		arrivalTime, err := time.Parse(time.RFC3339, flight.Arrival.Time)
		if err != nil {
			return nil, err
		}

		aircraft := flight.Aircraft

		flight := model.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.FlightID, constant.ProviderGarudaID),
			Provider: constant.ProviderGaruda,

			Airline: model.Airline{
				Name: flight.Airline,
				Code: flight.AirlineCode,
			},

			FlightNumber: flight.FlightID,

			Departure: model.Departure{
				Airport:   flight.Departure.Airport,
				City:      flight.Departure.City,
				Datetime:  departureTime.Format(time.RFC3339),
				Timestamp: departureTime.Unix(),
			},

			Arrival: model.Arrival{
				Airport:   flight.Arrival.Airport,
				City:      flight.Arrival.City,
				Datetime:  arrivalTime.Format(time.RFC3339),
				Timestamp: arrivalTime.Unix(),
			},

			Duration: model.Duration{
				TotalMinutes: flight.DurationMinutes,
				Formatted:    util.FormatDuration(flight.DurationMinutes),
			},

			Stops: flight.Stops,

			Price: model.Price{
				Amount:   flight.Price.Amount,
				Currency: flight.Price.Currency,
			},

			AvailableSeats: flight.AvailableSeats,
			CabinClass:     flight.FareClass,
			Aircraft:       &aircraft,
			Amenities:      flight.Amenities,
			Baggage:        buildGarudaBaggage(flight),
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func buildGarudaBaggage(flight dto.GarudaFlight) model.Baggage {
	return model.Baggage{
		CarryOn: fmt.Sprintf("%d", flight.Baggage.CarryOn),
		Checked: fmt.Sprintf("%d", flight.Baggage.Checked),
	}
}

package mapper

import (
	"fmt"
	"strings"
	"time"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/model"
	"github.com/fandrien/book-cabin/util"
)

func MapAirAsia(resp dto.AirAsiaResponse) ([]model.Flight, error) {

	flights := make([]model.Flight, 0, len(resp.Flights))

	for _, flight := range resp.Flights {

		departure, err := util.ParseRFC3339(flight.DepartTime)
		if err != nil {
			return nil, err
		}

		arrival, err := util.ParseRFC3339(flight.ArriveTime)
		if err != nil {
			return nil, err
		}

		stops := len(flight.Stops)

		layovers := 0
		for _, stop := range flight.Stops {
			layovers += stop.WaitTimeMinute
		}

		durationMinutes := int(flight.DurationHour*60) + layovers

		flight := model.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.FlightCode, constant.ProviderAirAsiaID),
			Provider: constant.ProviderAirAsia,

			Airline: model.Airline{
				Name: flight.Airline,
				Code: "QZ",
			},

			FlightNumber: flight.FlightCode,

			Departure: model.Departure{
				Airport:   flight.FromAirport,
				Datetime:  departure.Format(time.RFC3339),
				Timestamp: departure.Unix(),
			},

			Arrival: model.Arrival{
				Airport:   flight.ToAirport,
				Datetime:  arrival.Format(time.RFC3339),
				Timestamp: arrival.Unix(),
			},

			Duration: model.Duration{
				TotalMinutes: durationMinutes,
				Formatted:    util.FormatDuration(durationMinutes),
			},

			Stops: stops,

			Price: model.Price{
				Amount:   flight.PriceIDR,
				Currency: "IDR",
			},

			AvailableSeats: flight.Seats,

			CabinClass: flight.CabinClass,

			Baggage: buildAirAsiaBaggage(flight),
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func buildAirAsiaBaggage(flight dto.AirAsiaFlight) model.Baggage {
	baggage := model.Baggage{}
	if strings.Contains(strings.ToLower(flight.BaggageNote), "cabin baggage only") {
		baggage.CarryOn = "Cabin Baggage Only"
	}
	if strings.Contains(strings.ToLower(flight.BaggageNote), "checked bags additional fee") {
		baggage.Checked = "Additional fee"
	}
	return baggage
}

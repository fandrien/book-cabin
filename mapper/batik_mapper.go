package mapper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/model"
	"github.com/fandrien/book-cabin/util"
)

func MapBatik(resp dto.BatikResponse) ([]model.Flight, error) {

	flights := make([]model.Flight, 0, len(resp.Results))

	for _, flight := range resp.Results {

		departure, err := util.ParseTimeOffset(flight.DepartureDateTime)
		if err != nil {
			return nil, err
		}

		arrival, err := util.ParseTimeOffset(flight.ArrivalDateTime)
		if err != nil {
			return nil, err
		}
		layovers := 0
		Stops := flight.NumberOfStops

		for _, conn := range flight.Connections {
			stopDuration := conn.StopDuration[:len(conn.StopDuration)-1]
			num, err := strconv.Atoi(stopDuration)
			if err != nil {
				return nil, err
			}
			layovers += num
		}

		duration := int(arrival.Sub(departure).Minutes()) + layovers

		aircraft := flight.AircraftModel

		flight := model.Flight{
			ID:           fmt.Sprintf("%s_%s", flight.FlightNumber, constant.ProviderBatikID),
			Provider:     constant.ProviderBatik,
			FlightNumber: flight.FlightNumber,

			Airline: model.Airline{
				Name: flight.AirlineName,
				Code: flight.AirlineIATA,
			},

			Departure: model.Departure{
				Airport:   flight.Origin,
				Datetime:  departure.Format(time.RFC3339),
				Timestamp: departure.Unix(),
			},

			Arrival: model.Arrival{
				Airport:   flight.Destination,
				Datetime:  arrival.Format(time.RFC3339),
				Timestamp: arrival.Unix(),
			},

			Duration: model.Duration{
				TotalMinutes: duration,
				Formatted:    util.FormatDuration(duration),
			},

			Stops: Stops,

			Price: model.Price{
				Amount:   flight.Fare.TotalPrice,
				Currency: flight.Fare.Currency,
			},

			AvailableSeats: flight.SeatsAvailable,

			CabinClass: flight.Fare.Class,

			Aircraft: &aircraft,

			Baggage: buildBatikBaggage(flight),

			Amenities: append([]string{}, flight.OnboardServices...),
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func buildBatikBaggage(flight dto.BatikFlight) model.Baggage {
	baggage := model.Baggage{
		CarryOn: flight.BaggageInfo,
	}
	return baggage
}

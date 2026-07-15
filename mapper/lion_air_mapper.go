package mapper

import (
	"fmt"
	"time"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/model"
	"github.com/fandrien/book-cabin/util"
)

func MapLion(resp dto.LionResponse) ([]model.Flight, error) {

	flights := make([]model.Flight, 0, len(resp.Data.AvailableFlights))

	for _, flight := range resp.Data.AvailableFlights {

		departure, err := util.ParseTimeWithTimezone(
			flight.Schedule.Departure,
			flight.Schedule.DepartureTimezone,
		)
		if err != nil {
			return nil, err
		}

		arrival, err := util.ParseTimeWithTimezone(
			flight.Schedule.Arrival,
			flight.Schedule.ArrivalTimezone,
		)
		if err != nil {
			return nil, err
		}

		aircraft := flight.PlaneType

		stops := flight.StopCount
		layovers := 0
		for _, stop := range flight.Layovers {
			layovers += stop.DurationMinutes
		}

		flight := model.Flight{
			ID:       fmt.Sprintf("%s_%s", flight.ID, constant.ProviderLionID),
			Provider: constant.ProviderLion,

			Airline: model.Airline{
				Name: flight.Carrier.Name,
				Code: flight.Carrier.IATA,
			},

			FlightNumber: flight.ID,

			Departure: model.Departure{
				Airport:   flight.Route.From.Code,
				City:      flight.Route.From.City,
				Datetime:  departure.Format(time.RFC3339),
				Timestamp: departure.Unix(),
			},

			Arrival: model.Arrival{
				Airport:   flight.Route.To.Code,
				City:      flight.Route.To.City,
				Datetime:  arrival.Format(time.RFC3339),
				Timestamp: arrival.Unix(),
			},

			Duration: model.Duration{
				TotalMinutes: flight.FlightTime + layovers,
				Formatted:    util.FormatDuration(flight.FlightTime),
			},

			Stops: stops,

			Price: model.Price{
				Amount:   flight.Pricing.Total,
				Currency: flight.Pricing.Currency,
			},

			AvailableSeats: flight.SeatsLeft,

			CabinClass: flight.Pricing.FareType,

			Aircraft: &aircraft,

			Baggage: buildLionBaggage(flight),

			Amenities: buildLionAmenities(flight),
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func buildLionBaggage(flight dto.LionFlight) model.Baggage {
	baggage := model.Baggage{
		CarryOn: flight.Services.BaggageAllowance.Cabin,
		Checked: flight.Services.BaggageAllowance.Hold,
	}
	return baggage
}

func buildLionAmenities(flight dto.LionFlight) []string {
	amenities := make([]string, 0)
	if flight.Services.WifiAvailable {
		amenities = append(amenities, "wifi")
	}
	if flight.Services.MealsIncluded {
		amenities = append(amenities, "meals")
	}
	return amenities
}

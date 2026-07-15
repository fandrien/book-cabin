package dto

type AirAsiaResponse struct {
	Status  string          `json:"status"`
	Flights []AirAsiaFlight `json:"flights"`
}

type AirAsiaFlight struct {
	FlightCode   string        `json:"flight_code"`
	Airline      string        `json:"airline"`
	FromAirport  string        `json:"from_airport"`
	ToAirport    string        `json:"to_airport"`
	DepartTime   string        `json:"depart_time"`
	ArriveTime   string        `json:"arrive_time"`
	DurationHour float64       `json:"duration_hours"`
	DirectFlight bool          `json:"direct_flight"`
	Stops        []AirAsiaStop `json:"stops,omitempty"`
	PriceIDR     int           `json:"price_idr"`
	Seats        int           `json:"seats"`
	CabinClass   string        `json:"cabin_class"`
	BaggageNote  string        `json:"baggage_note"`
}

type AirAsiaStop struct {
	Airport        string `json:"airport"`
	WaitTimeMinute int    `json:"wait_time_minutes"`
}

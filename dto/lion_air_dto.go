package dto

type LionResponse struct {
	Success bool     `json:"success"`
	Data    LionData `json:"data"`
}

type LionData struct {
	AvailableFlights []LionFlight `json:"available_flights"`
}

type LionFlight struct {
	ID         string        `json:"id"`
	Carrier    LionCarrier   `json:"carrier"`
	Route      LionRoute     `json:"route"`
	Schedule   LionSchedule  `json:"schedule"`
	FlightTime int           `json:"flight_time"`
	IsDirect   bool          `json:"is_direct"`
	StopCount  int           `json:"stop_count,omitempty"`
	Layovers   []LionLayover `json:"layovers,omitempty"`
	Pricing    LionPricing   `json:"pricing"`
	SeatsLeft  int           `json:"seats_left"`
	PlaneType  string        `json:"plane_type"`
	Services   LionServices  `json:"services"`
}

type LionCarrier struct {
	Name string `json:"name"`
	IATA string `json:"iata"`
}

type LionRoute struct {
	From LionAirport `json:"from"`
	To   LionAirport `json:"to"`
}

type LionAirport struct {
	Code string `json:"code"`
	Name string `json:"name"`
	City string `json:"city"`
}

type LionSchedule struct {
	Departure         string `json:"departure"`
	DepartureTimezone string `json:"departure_timezone"`
	Arrival           string `json:"arrival"`
	ArrivalTimezone   string `json:"arrival_timezone"`
}

type LionPricing struct {
	Total    int    `json:"total"`
	Currency string `json:"currency"`
	FareType string `json:"fare_type"`
}

type LionServices struct {
	WifiAvailable    bool                 `json:"wifi_available"`
	MealsIncluded    bool                 `json:"meals_included"`
	BaggageAllowance LionBaggageAllowance `json:"baggage_allowance"`
}

type LionBaggageAllowance struct {
	Cabin string `json:"cabin"`
	Hold  string `json:"hold"`
}

type LionLayover struct {
	Airport         string `json:"airport"`
	DurationMinutes int    `json:"duration_minutes"`
}

package model

type SearchRequest struct {
	Origin        string   `json:"origin"`
	Destination   string   `json:"destination"`
	DepartureDate string   `json:"departureDate"`
	ReturnDate    *string  `json:"returnDate,omitempty"`
	Passengers    int      `json:"passengers"`
	CabinClass    string   `json:"cabinClass"`
	ArrivalDate   string   `json:"arrivalDate"`
	Airlines      []string `json:"airlines"`
	Stops         *int     `json:"stops"`
	MinDuration   *int     `json:"minDuration"`
	MaxDuration   *int     `json:"maxDuration"`
	MinPrice      *int     `json:"minPrice"`
	MaxPrice      *int     `json:"maxPrice"`
	SortBy        string   `json:"sort_by"`    // price, duration, departure, arrival
	SortOrder     string   `json:"sort_order"` // asc, desc
}

type SearchResponse struct {
	TotalFlights int                `json:"total_flights"`
	SearchTimeMs int64              `json:"search_time_ms"`
	ProviderLogs []ProviderResult   `json:"provider_logs"`
	Flights      []Flight           `json:"flights"`
	Comparisons  []FlightComparison `json:"comparisons,omitempty"`
}

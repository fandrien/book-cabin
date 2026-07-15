package model

type SearchRequest struct {
	Origin        string  `json:"origin"`
	Destination   string  `json:"destination"`
	DepartureDate string  `json:"departureDate"`
	ReturnDate    *string `json:"returnDate,omitempty"`
	Passengers    int     `json:"passengers"`
	CabinClass    string  `json:"cabinClass"`
	ArrivalDate   string  `json:"arrivalDate"`
	Airline       string  `json:"airline"`
	Stop          *int    `json:"stop"`
	Duration      *int    `json:"duration"`
	Price         *int    `json:"price"`
	SortBy        string  `json:"sort_by"`    // price, duration, departure, arrival
	SortOrder     string  `json:"sort_order"` // asc, desc
}

type SearchResponse struct {
	Total        int              `json:"total"`
	SearchTimeMs int64            `json:"search_time_ms"`
	Providers    []ProviderResult `json:"providers"`
	Flights      []Flight         `json:"flights"`
}

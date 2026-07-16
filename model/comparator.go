package model

type Offer struct {
	FlightNumber string `json:"flight_number"`
	CabinClass   string `json:"cabin_class"`
	Provider     string `json:"provider"`
	Price        int    `json:"price"`
}

type FlightComparison struct {
	Key       string  `json:"key"`
	BestPrice int     `json:"best_price"`
	Savings   int     `json:"savings"`
	Offers    []Offer `json:"offers"`
}

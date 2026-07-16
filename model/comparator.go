package model

type Offer struct {
	Provider string `json:"provider"`
	Price    int    `json:"price"`
}

type FlightComparison struct {
	Key       string  `json:"key"`
	BestPrice int     `json:"best_price"`
	Savings   int     `json:"savings"`
	Offers    []Offer `json:"offers"`
}

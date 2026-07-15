package model

type ProviderResult struct {
	Name        string `json:"name"`
	Success     bool   `json:"success"`
	FlightCount int    `json:"flight_count"`
	DurationMs  int64  `json:"duration_ms"`
	Error       string `json:"error,omitempty"`
}

package constant

import "time"

const ProviderTimeout = 500 * time.Millisecond

const (
	ProviderGarudaID  = "Garuda_Indonesia"
	ProviderLionID    = "Lion_Air"
	ProviderBatikID   = "Batik_Air"
	ProviderAirAsiaID = "AirAsia"
)

const (
	ProviderGaruda  = "Garuda Indonesia"
	ProviderLion    = "Lion Air"
	ProviderBatik   = "Batik Air"
	ProviderAirAsia = "AirAsia"
)

const (
	GarudaDataPath  = "external/garuda_indonesia_search_response.json"
	LionDataPath    = "external/lion_air_search_response.json"
	BatikDataPath   = "external/batik_air_search_response.json"
	AirAsiaDataPath = "external/airasia_search_response.json"
)

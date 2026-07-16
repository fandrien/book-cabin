package provider

import "golang.org/x/time/rate"

var (
	GarudaLimiter  = rate.NewLimiter(5, 10)
	LionLimiter    = rate.NewLimiter(5, 10)
	AirAsiaLimiter = rate.NewLimiter(5, 10)
	BatikLimiter   = rate.NewLimiter(5, 10)
)

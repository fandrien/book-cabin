package provider

import (
	"context"
	"errors"
	"math/rand"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/helper"
	"github.com/fandrien/book-cabin/loader"
	"github.com/fandrien/book-cabin/mapper"
	"github.com/fandrien/book-cabin/model"
)

type AirAsiaProvider struct{}

func NewAirAsiaProvider() *AirAsiaProvider {
	return &AirAsiaProvider{}
}

func (g *AirAsiaProvider) Name() string {
	return constant.ProviderAirAsia
}

func (g *AirAsiaProvider) Search(
	ctx context.Context,
) ([]model.Flight, error) {

	var flights []model.Flight

	// Retry with exponential backoff
	err := helper.Retry(
		ctx,
		3,
		func() error {

			result, err := g.search(ctx)
			if err != nil {
				return err
			}

			flights = result
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return flights, nil
}

func (g *AirAsiaProvider) search(
	ctx context.Context,
) ([]model.Flight, error) {

	// Set Limiter
	if err := AirAsiaLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	// AirAsia delay 50-150ms
	if err := loader.RandomDelay(ctx, 50, 150); err != nil {
		return nil, err
	}

	// 90% success rate
	if rand.Intn(100) >= 90 {
		return nil, errors.New("airasia provider unavailable")
	}

	resp, err := loader.LoadJSON[dto.AirAsiaResponse](
		constant.AirAsiaDataPath,
	)
	if err != nil {
		return nil, err
	}

	return mapper.MapAirAsia(resp)
}

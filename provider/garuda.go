package provider

import (
	"context"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/helper"
	"github.com/fandrien/book-cabin/loader"
	"github.com/fandrien/book-cabin/mapper"
	"github.com/fandrien/book-cabin/model"
)

type GarudaProvider struct{}

func NewGarudaProvider() *GarudaProvider {
	return &GarudaProvider{}
}

func (g *GarudaProvider) Name() string {
	return constant.ProviderGaruda
}

func (g *GarudaProvider) Search(ctx context.Context) ([]model.Flight, error) {
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

func (g *GarudaProvider) search(ctx context.Context) ([]model.Flight, error) {
	// Set Limiter
	if err := GarudaLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	// Garuda delay 100-200ms
	if err := loader.RandomDelay(ctx, 100, 200); err != nil {
		return nil, err
	}

	// Read Source Data (JSON file)
	resp, err := loader.LoadJSON[dto.GarudaResponse](
		constant.GarudaDataPath,
	)
	if err != nil {
		return nil, err
	}

	// Normalize Data
	return mapper.MapGaruda(resp)
}

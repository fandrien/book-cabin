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

type BatikProvider struct{}

func NewBatikProvider() *BatikProvider {
	return &BatikProvider{}
}

func (g *BatikProvider) Name() string {
	return constant.ProviderBatik
}

func (g *BatikProvider) Search(
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

func (g *BatikProvider) search(
	ctx context.Context,
) ([]model.Flight, error) {

	// Set Limiter
	if err := BatikLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	// Batik delay 200-400ms
	if err := loader.RandomDelay(ctx, 200, 400); err != nil {
		return nil, err
	}

	// Read Source Data (JSON file)
	resp, err := loader.LoadJSON[dto.BatikResponse](
		constant.BatikDataPath,
	)
	if err != nil {
		return nil, err
	}

	// Normalize Data
	return mapper.MapBatik(resp)
}

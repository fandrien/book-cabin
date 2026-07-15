package provider

import (
	"context"
	"errors"
	"math/rand"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
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
	req model.SearchRequest,
) ([]model.Flight, error) {

	//AirAsia delay 50-150ms & occasionaly fails
	if err := loader.RandomDelay(ctx, 50, 150); err != nil {
		return nil, err
	}

	// 90% success rate
	if rand.Intn(100) >= 90 {
		return nil, errors.New("airasia provider unavailable")
	}

	//Read Source Data (JSON file)
	resp, err := loader.LoadJSON[dto.AirAsiaResponse](constant.AirAsiaDataPath)
	if err != nil {
		return nil, err
	}

	//Normalize Data
	return mapper.MapAirAsia(resp)
}

package provider

import (
	"context"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
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
	req model.SearchRequest,
) ([]model.Flight, error) {

	if err := BatikLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	//Batik delay 200-400ms
	if err := loader.RandomDelay(ctx, 200, 400); err != nil {
		return nil, err
	}

	//Read Source Data (JSON file)
	resp, err := loader.LoadJSON[dto.BatikResponse](constant.BatikDataPath)
	if err != nil {
		return nil, err
	}

	//Normalize Data
	return mapper.MapBatik(resp)
}

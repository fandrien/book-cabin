package provider

import (
	"context"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/dto"
	"github.com/fandrien/book-cabin/loader"
	"github.com/fandrien/book-cabin/mapper"
	"github.com/fandrien/book-cabin/model"
)

type LionProvider struct{}

func NewLionProvider() *LionProvider {
	return &LionProvider{}
}

func (l *LionProvider) Name() string {
	return constant.ProviderLion
}

func (l *LionProvider) Search(
	ctx context.Context,
	req model.SearchRequest,
) ([]model.Flight, error) {

	//Lion Air delay 100-200ms
	if err := loader.RandomDelay(ctx, 100, 200); err != nil {
		return nil, err
	}

	//Read Source Data (JSON file)
	resp, err := loader.LoadJSON[dto.LionResponse](constant.LionDataPath)
	if err != nil {
		return nil, err
	}

	//Normalize Data
	return mapper.MapLion(resp)
}

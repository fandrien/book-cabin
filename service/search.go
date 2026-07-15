package service

import (
	"context"
	"time"

	"github.com/fandrien/book-cabin/aggregation"
	"github.com/fandrien/book-cabin/cache"
	"github.com/fandrien/book-cabin/model"
)

type SearchService struct {
	aggregator *aggregation.Aggregator
	cache      cache.Cache
}

func NewSearchService(
	aggregator *aggregation.Aggregator,
	cache cache.Cache,
) *SearchService {

	return &SearchService{
		aggregator: aggregator,
		cache:      cache,
	}
}

func (s *SearchService) Search(
	ctx context.Context,
	req model.SearchRequest,
) (*model.SearchResponse, error) {

	//get cache
	cacheKey := cache.BuildKey(req)
	if response, found := s.cache.Get(cacheKey); found {
		return response, nil
	}

	start := time.Now()

	result := s.aggregator.Aggregate(
		ctx,
		req,
	)

	flights := filterFlights(
		req,
		result.Flights,
	)

	sortFlights(
		req,
		flights,
	)

	response := &model.SearchResponse{

		Total: len(flights),

		Providers: result.Providers,

		SearchTimeMs: time.Since(start).Milliseconds(),

		Flights: flights,
	}

	s.cache.Set(cacheKey, response)

	return response, nil
}

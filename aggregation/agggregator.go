package aggregation

import (
	"context"
	"sync"
	"time"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/model"
	"github.com/fandrien/book-cabin/provider"
)

type Aggregator struct {
	providers []provider.Provider
}

func New(
	providers []provider.Provider,
) *Aggregator {

	return &Aggregator{
		providers: providers,
	}
}

func (a *Aggregator) Aggregate(
	ctx context.Context,
	req model.SearchRequest,
) Result {
	var (
		wg sync.WaitGroup
		mu sync.Mutex

		result Result
	)

	for _, p := range a.providers {

		wg.Add(1)

		go func(provider provider.Provider) {

			defer wg.Done()

			providerCtx, cancel := context.WithTimeout(
				ctx,
				constant.ProviderTimeout,
			)
			defer cancel()

			start := time.Now()

			flights, err := provider.Search(providerCtx)

			stat := model.ProviderResult{
				Name:       provider.Name(),
				DurationMs: time.Since(start).Milliseconds(),
			}

			if err != nil {
				stat.Success = false
				stat.Error = err.Error()

			} else {
				stat.Success = true
				stat.FlightCount = len(flights)
			}

			mu.Lock()

			result.Providers = append(
				result.Providers,
				stat,
			)

			result.Flights = append(
				result.Flights,
				flights...,
			)

			mu.Unlock()

		}(p)

	}

	wg.Wait()

	return result
}

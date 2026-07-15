package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fandrien/book-cabin/aggregation"
	"github.com/fandrien/book-cabin/cache"
	"github.com/fandrien/book-cabin/handler"
	"github.com/fandrien/book-cabin/provider"
	"github.com/fandrien/book-cabin/router"
	"github.com/fandrien/book-cabin/service"
)

func main() {

	// Register providers
	providers := []provider.Provider{
		provider.NewGarudaProvider(),
		provider.NewLionProvider(),
		provider.NewAirAsiaProvider(),
		provider.NewBatikProvider(),
	}

	// Create Aggregator
	aggregator := aggregation.New(providers)

	// Create Memory Cache
	memoryCache := cache.NewMemoryCache(5 * time.Minute)

	// Create Service
	searchService := service.NewSearchService(aggregator, memoryCache)

	// Create Handler
	searchHandler := handler.NewSearchHandler(searchService)

	// Setup Router
	r := router.NewRouter(searchHandler)

	log.Println("Server started at :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}

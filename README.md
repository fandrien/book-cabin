# Book Cabin

A Go-based flight search aggregator that collects flight data concurrently from multiple airline providers, normalizes the responses into a unified model, and returns the best available results with filtering, sorting and caching.

## Features

* Concurrent search across multiple airline providers
* Provider pattern for easy extensibility
* DTO → Domain Model mapping
* Generic JSON loader
* Aggregation layer
* In-memory cache
* Filtering
* Sorting
* Best Value ranking algorithm
* Request validation
* Partial failure handling
* Context propagation
* Provider timeout
* Provider rate limiting

---

# Supported Providers

* Garuda Indonesia
* Lion Air
* AirAsia
* Batik Air

Each provider has:

* DTO
* Mapper
* Provider implementation
* Mock JSON response

---

# Project Structure

```
book-cabin/
│
├── api/
│   └── main.go
│
├── aggregation/
├── cache/
├── constant/
├── dto/
├── external/
├── handler/
├── loader/
├── mapper/
├── model/
├── provider/
├── response/
├── router/
├── service/
├── util/
├── validator/
│
├── go.mod
└── README.md
```

---

# Architecture

```
                HTTP Request
                     │
                     ▼
                Search Handler
                     │
                     ▼
             Request Validation
                     │
                     ▼
              Search Service
                     │
          ┌──────────┴──────────┐
          │                     │
          ▼                     ▼
      Cache Lookup         Aggregator
                                  │
          ┌───────────────────────┼──────────────────────┐
          ▼                       ▼                      ▼
      Garuda Provider      Lion Provider       AirAsia Provider
                                  │
                             Batik Provider
                                  │
                                  ▼
                          Unified Flight Model
                                  │
                                  ▼
                     Filtering → Sorting → Response
```

---

# Setup

## Prerequisites

* Go 1.26+
* Git

Clone the repository

```bash
git clone https://github.com/fandrien/book-cabin.git

cd book-cabin
```

Install dependencies

```bash
go mod tidy
```

Run the application

```bash
go run ./api
```

Server will start on

```
http://localhost:8080
```

---

# API

## Search Flights

```
POST /search
```

### Request

```json
{
  "origin": "CGK",
  "destination": "DPS",
  "departureDate": "2025-12-15",
  "returnDate": "",
  "arrivalDate": "2025-12-15",
  "airlines": ["AirAsia", "Lion Air"],
  "stops": 1,
  "minDuration": 60,
  "maxDuration": 500,
  "minPrice": 100000,
  "maxPrice": 1000000,
  "sort_by": "best",
  "sort_order": "asc"
}
```

---

# Available Filters

| Filter           | Description         |
| ---------------- | ------------------- |
| origin           | Origin airport      |
| destination      | Destination airport |
| departureDate    | Departure date      |
| arrivalDate      | Arrival date        |
| returnDate       | Return date         |
| airlines         | Filter by airlines  |
| minPrice         | Min Price           |
| maxPrice         | Max Price           |
| minDuration      | Min Flight Duration |
| maxDuration      | Max Flight Duration |
| stops            | Number of stops     |


---

# Sorting

Supported sorting fields

* best (default)
* price
* duration
* departure
* arrival

Supported order

* asc
* desc

---

## Best Value Algorithm

If no sorting option is specified, flights are ranked using a custom Best Value algorithm.

The score combines:

- Ticket price (lower is better)
- Total travel duration (shorter is better)
- Number of stops (fewer stops are strongly preferred)

Flights with fewer stops receive a significant advantage, making direct flights rank higher whenever possible.

The flight with the lowest score is considered the best value.

---

# Caching

Search results are cached in memory.

Cache key is generated from the search request parameters.

Benefits:

* Faster repeated searches
* Reduced provider calls
* Lower response time

---

# Timeout

Each provider request is executed with its own timeout using Go Context.

If a provider exceeds the configured timeout:

* The provider is marked as failed.
* Other providers continue processing.
* Partial results are still returned.

---

# Partial Failure

The system tolerates provider failures.

Example:

```
Garuda     ✅

Lion       ✅

AirAsia    ❌ Timeout

Batik      ✅
```

The API still returns available flights from successful providers.

---

# Validation

The following validations are performed:

* Origin is required
* Destination is required
* Origin cannot equal destination
* Departure date format validation
* Return date format validation
* Return date cannot be earlier than departure date

---

# Concurrency

Provider searches run concurrently using:

* Goroutines
* sync.WaitGroup
* sync.Mutex

This reduces overall search latency.

---

# Provider Rate Limiting

To prevent overwhelming external provider APIs and to respect
third-party quotas, each provider is protected by its own
token-bucket rate limiter.

Example:
- Garuda: 5 requests/second
- Lion Air: 5 requests/second
- AirAsia: 5 requests/second
- Batik Air: 5 requests/second

---

# Technologies

* Go
* net/http
* Context
* Goroutines
* WaitGroup
* Mutex
* JSON
* Generic Functions

---




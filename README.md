# Book Cabin

A Go-based flight search aggregator that collects flight data concurrently from multiple airline providers, normalizes the responses into a unified model, and returns the best available results with filtering, ing, and caching.

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
go run .
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
  "departure_date": "2025-12-15",
  "returnDate": "2025-12-18",
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
| departure_date   | Departure date      |
| return_date      | Return date         |
| airline          | Filter by airline   |
| price            | Price               |
| duration         | Duration            |
| stops            | Number of stops     |


---

# ing

Supported ing fields

* best (default)
* price
* duration
* departure
* arrival

Supported order

* asc
* desc

---

# Best Value Algorithm

If no ing option is specified, flights are ranked using a custom Best Value algorithm.

The score considers:

* Direct flights are preferred
* Lower ticket prices
* Shorter total travel duration (including layovers)

This provides a balance between affordability and convenience.

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




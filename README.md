# exchangerateapi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/exchangerateapinet/exchangerateapi-go.svg)](https://pkg.go.dev/github.com/exchangerateapinet/exchangerateapi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/exchangerateapinet/exchangerateapi-go)](https://goreportcard.com/report/github.com/exchangerateapinet/exchangerateapi-go)
[![GitHub release](https://img.shields.io/github/v/release/exchangerateapinet/exchangerateapi-go?display_name=tag&sort=semver)](https://github.com/exchangerateapinet/exchangerateapi-go/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Straightforward Go client for exchangerateapi.net with minimal dependencies.

- Website: [exchangerateapi.net](https://exchangerateapi.net)

## Install

```bash
go get github.com/exchangerateapinet/exchangerateapi-go@latest
```

## Quick start

```go
import api "github.com/exchangerateapinet/exchangerateapi-go/exchangerateapi"
client := api.New("YOUR_API_KEY")
```

## Usage

```go
import (
    api "github.com/exchangerateapinet/exchangerateapi-go/exchangerateapi"
)

func example(client *api.Client) {
    // Latest with base only
    latest, err := client.Latest("USD", nil)
    if err != nil { /* handle */ }

    // Latest with a subset of symbols
    subset, err := client.Latest("EUR", []string{"USD", "GBP", "JPY"})
    if err != nil { /* handle */ }

    // Historical by date
    hist, err := client.Historical("2024-01-02", "USD", nil)
    if err != nil { /* handle */ }

    _, _, _ = latest, subset, hist
}
```

## Configuration

- API key is required; obtain one at [exchangerateapi.net](https://exchangerateapi.net).
- The client uses `http.DefaultClient`. You can set a custom client with timeouts or retry middleware by assigning `Client.HTTP` after `New(...)`.
- The base URL defaults to `https://api.exchangerateapi.net/v1`.

```go
c := api.New(os.Getenv("EXCHANGERATEAPI_KEY"))
c.HTTP = &http.Client{ Timeout: 10 * time.Second }
// c.BaseURL = "https://api.exchangerateapi.net/v1" // override if needed
```

## Errors

All methods return a Go error if the request fails or the API returns an error payload. Check the error value and decide whether to retry or surface it to callers.

## Capabilities

- Fetch the newest rates for a base currency (optional symbol filter)
- Retrieve rates for a historical date (optional symbol filter)

### Examples

Run with an API key in your environment:

```bash
EXCHANGERATEAPI_KEY=your_api_key go run ./examples/latest
EXCHANGERATEAPI_KEY=your_api_key go run ./examples/historical
```

## Free usage

A free tier is available for light usage and development. It requires an API key and provides basic access to the latest and historical endpoints with rate limits. Limits and details may change—please see the pricing/plan information on [exchangerateapi.net](https://exchangerateapi.net) for the most current terms.

## License

MIT

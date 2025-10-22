# exchangerateapi-go

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

## Capabilities

- Fetch the newest rates for a base currency (optional symbol filter)
- Retrieve rates for a historical date (optional symbol filter)

### Examples

Run with an API key in your environment:

```bash
EXCHANGERATEAPI_KEY=your_api_key go run ./examples/latest
EXCHANGERATEAPI_KEY=your_api_key go run ./examples/historical
```

## License

MIT

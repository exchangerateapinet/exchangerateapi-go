package exchangerateapi

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

// Client provides access to exchangerateapi.net endpoints.
type Client struct {
	APIKey  string
	BaseURL string
	HTTP    *http.Client
}

// New creates a new Client.
func New(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.exchangerateapi.net/v1",
		HTTP:    http.DefaultClient,
	}
}

type latestResponse struct {
	Rates map[string]float64 `json:"rates"`
	Error any               `json:"error"`
}

type historicalResponse struct {
	Rates map[string]float64 `json:"rates"`
	Error any               `json:"error"`
}

// Latest retrieves the newest conversion rates for a base currency.
func (c *Client) Latest(base string, symbols []string) (map[string]float64, error) {
	if base == "" {
		return nil, errors.New("base is required")
	}
	q := url.Values{}
	q.Set("base", base)
	q.Set("apikey", c.APIKey)
	if len(symbols) > 0 {
		q.Set("symbols", joinSymbols(symbols))
	}
	endpoint := fmt.Sprintf("%s/latest?%s", c.BaseURL, q.Encode())
	var r latestResponse
	if err := c.getJSON(endpoint, &r); err != nil {
		return nil, err
	}
	if r.Error != nil {
		return nil, fmt.Errorf("api error: %v", r.Error)
	}
	return r.Rates, nil
}

// Historical retrieves rates for a given date and base currency.
func (c *Client) Historical(date, base string, symbols []string) (map[string]float64, error) {
	if date == "" || base == "" {
		return nil, errors.New("date and base are required")
	}
	q := url.Values{}
	q.Set("date", date)
	q.Set("base", base)
	q.Set("apikey", c.APIKey)
	if len(symbols) > 0 {
		q.Set("symbols", joinSymbols(symbols))
	}
	endpoint := fmt.Sprintf("%s/historical?%s", c.BaseURL, q.Encode())
	var r historicalResponse
	if err := c.getJSON(endpoint, &r); err != nil {
		return nil, err
	}
	if r.Error != nil {
		return nil, fmt.Errorf("api error: %v", r.Error)
	}
	return r.Rates, nil
}

func (c *Client) getJSON(endpoint string, out any) error {
	resp, err := c.HTTP.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(out)
}

func joinSymbols(s []string) string {
    // Join with commas; url.Values in caller encodes when added to the query.
    return strings.Join(s, ",")
}

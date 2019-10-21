package pokeclient

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

// Client represents the client state for the API
type Client struct {
	RestyClient *resty.Client
}

// Config defines configuration behavior for the Pokemon API
type Config struct {
	APIKey     string
	BaseURL    string
	HostHeader string
	Debug      bool
}

// NewClient creates a new instance of the HTTP using the provided configuration
func NewClient(config Config) Client {
	r := resty.New()

	baseURL := config.BaseURL

	if baseURL == "" {
		baseURL = "https://pokemon-go1.p.rapidapi.com"
	}

	hostHeader := config.HostHeader

	if hostHeader == "" {
		hostHeader = "pokemon-go1.p.rapidapi.com"
	}

	r.SetHeader("x-rapidapi-key", config.APIKey)
	r.SetHostURL(baseURL)
	r.SetHeader("x-rapidapi-host", hostHeader)

	if config.Debug {
		r.SetDebug(true)
	}

	c := Client{
		RestyClient: r,
	}

	return c
}

// Do executes the requested HTTP method for the path
func (c *Client) Do(method string, path string) ([]byte, error) {
	r := c.RestyClient.R().SetHeader("Content-Type", "application/json")

	apiResponse, err := r.Execute(method, path)

	if err != nil {
		return nil, err
	}

	body := apiResponse.Body()

	statusClass := apiResponse.StatusCode() / 100 % 10

	if statusClass == 2 {
		return body, nil
	}

	return nil, fmt.Errorf("Returned an error of type %d", statusClass)
}

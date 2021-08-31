package http

import (
	"net/http"
)

// client is a wrapper around http.Client
type Client struct {
	client *http.Client
}

// New returns a new instance of client which implements Client interface
func New() Client {
	return Client{client: http.DefaultClient}
}

// Get sends a GET request
func (c Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

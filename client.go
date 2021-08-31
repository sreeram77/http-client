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
func (c Client) Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key := range headers {
		req.Header.Add(key, headers[key])
	}

	return c.client.Do(req)
}

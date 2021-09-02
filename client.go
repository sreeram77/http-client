package http

import (
	"fmt"
	"net/http"
	"net/url"
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
func (c Client) Get(u string, headers map[string]string, params map[string]interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	for key := range headers {
		req.Header.Add(key, headers[key])
	}

	q := url.Values{}

	for key, value := range params {
		switch t := value.(type) {
		case []string:
			for i := range t {
				q.Set(key, t[i])
			}
		case string:
			q.Set(key, t)
		default:
			q.Set(key, fmt.Sprintf("%v", value))

		}
	}

	req.URL.RawQuery = q.Encode()

	return c.client.Do(req)
}

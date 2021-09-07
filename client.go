package http

import (
	"bytes"
	"fmt"
	"net/http"
	netURL "net/url"
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
func (c Client) Get(url string, headers map[string]string, params map[string]interface{}) (*http.Response, error) {
	req, err := generateRequest(url, http.MethodPost, headers, params, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Post sends a POST request
func (c Client) Post(url string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Response, error) {
	req, err := generateRequest(url, http.MethodGet, headers, params, body)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func generateRequest(url, method string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key := range headers {
		req.Header.Add(key, headers[key])
	}

	q := netURL.Values{}

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

	return req, nil
}

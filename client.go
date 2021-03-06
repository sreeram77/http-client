package http

import (
	"bytes"
	"fmt"
	"net/http"
	netURL "net/url"
	"time"
)

// client is a wrapper around http.Client
type Client struct {
	client *http.Client
}

// New returns a new instance of client which implements Client interface
func New() Client {
	return Client{client: http.DefaultClient}
}

// NewClientWithTimeout returns an instance of Client with a Timeout
// Timeout for DefaultClient is 0, which means the connection never times out
func NewClientWithTimeout(t time.Duration) Client {
	c := http.DefaultClient
	c.Timeout = t

	return Client{client: c}
}

// Get sends a GET request
func (c Client) Get(url string, headers map[string]string, params map[string]interface{}) (*http.Response, error) {
	req, err := GenerateRequest(url, http.MethodGet, headers, params, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Post sends a POST request
func (c Client) Post(url string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Response, error) {
	req, err := GenerateRequest(url, http.MethodPost, headers, params, body)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Delete sends a DELETE request
func (c Client) Delete(url string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Response, error) {
	req, err := GenerateRequest(url, http.MethodDelete, headers, params, body)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Put sends a PUT request
func (c Client) Put(url string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Response, error) {
	req, err := GenerateRequest(url, http.MethodPut, headers, params, body)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Patch sends a PATCH request
func (c Client) Patch(url string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Response, error) {
	req, err := GenerateRequest(url, http.MethodGet, headers, params, body)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// GenerateRequest creates HTTP Request from input params
func GenerateRequest(url, method string, headers map[string]string, params map[string]interface{}, body []byte) (*http.Request, error) {
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

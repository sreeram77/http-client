package http

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_Get(t *testing.T) {
	type testCase struct {
		id       int
		input    string
		response *http.Response
	}

	testCases := []testCase{
		{
			id:    1,
			input: "http://google.com",
			response: &http.Response{
				StatusCode: http.StatusOK,
			},
		},
	}

	for _, tc := range testCases {
		c := New()
		resp, err := c.Get(tc.input, nil, nil)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(resp.StatusCode, tc.response.StatusCode) {
			t.Errorf("testcase: %v Expected Output: %v Actual Output: %v", tc.id, tc.response, resp)
		}
	}
}

func Test_GenerateRequest(t *testing.T) {
	type testCase struct {
		id     int
		url    string
		method string
		header map[string]string
		body   []byte
		param  map[string]interface{}
	}

	testCases := []testCase{
		{
			id:     1,
			url:    "https://google.com",
			method: http.MethodGet,
			header: nil,
			body:   nil,
			param:  nil,
		},
	}

	for _, tc := range testCases {
		req, err := generateRequest(tc.url, tc.method, tc.header, tc.param, tc.body)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(req.Method, tc.method) {
			t.Errorf("testcase: %v Expected Output: %v Actual Output: %v", tc.id, tc.method, req.Method)
		}
	}
}

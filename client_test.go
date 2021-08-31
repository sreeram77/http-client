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
		err      error
	}

	testCases := []testCase{
		{
			id:    1,
			input: "http://google.com",
			response: &http.Response{
				StatusCode: http.StatusOK,
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		c := New()
		resp, err := c.Get(tc.input, nil)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(resp.StatusCode, tc.response.StatusCode) {
			t.Errorf("testcase: %v Expected Output: %v Actual Output: %v", tc.id, tc.response, resp)
		}
	}
}

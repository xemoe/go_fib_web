package actions

import (
	"encoding/json"
	"net/http"
)

func (as *ActionSuite) Test_FibHandler() {
	var tests = []struct {
		num                string
		expectedStatusCode int
		expectedResult     string
	}{
		{"1", http.StatusOK, "1"},
		{"2", http.StatusOK, "2"},
		{"3", http.StatusOK, "fizz"},
		{"5", http.StatusOK, "buzz"},
		{"15", http.StatusOK, "fizzbuzz"},
	}

	for _, t := range tests {
		res := as.HTML("/api/v1/fib/" + t.num).Get()

		var response Message
		json.Unmarshal(res.Body.Bytes(), &response)

		as.Equal(t.expectedStatusCode, res.Code)
		as.Equal(t.expectedResult, response.Result)
	}
}

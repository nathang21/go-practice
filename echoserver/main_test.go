package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		path    string
		handler http.HandlerFunc
		err     bool
	}{
		{"/", hello, false},
		{"/test", hello, false},
		{"/hello", hello, false},
		{"", hello, false},
		{"//", hello, false},
		{"/errorz", errorz, true},
	}

	for _, test := range tests {
		path := fmt.Sprintf("%s", test.path)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(test.handler)
		handler.ServeHTTP(rr, req)

		// In this case, our MetricsHandler returns a non-200 response
		// for a route variable it doesn't know about.
		if rr.Code == http.StatusOK && test.err != false {
			t.Errorf("handler should have passed on route %s: got %v want %v",
				test.path, rr.Code, http.StatusOK)
		} else if rr.Code == http.StatusInternalServerError && test.err != true {
			t.Errorf("handler should have failed on route %s: got %v want %v",
				test.path, rr.Code, http.StatusInternalServerError)
		}
	}
}

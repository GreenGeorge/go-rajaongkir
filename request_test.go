package rajaongkir

import (
	"net/http"
	"testing"
)

func TestCreateTargetURL(t *testing.T) {
	ro := New("APIKEY12345", "api.rajaongkir.com/starter", nil)

	tables := []struct {
		endpoint string
		expected string
	}{
		{"/city", "https://api.rajaongkir.com/starter/city"},
		{"/province", "https://api.rajaongkir.com/starter/province"},
		{"/cost", "https://api.rajaongkir.com/starter/cost"},
	}

	for _, table := range tables {
		result := ro.createTargetURL(table.endpoint)
		if result != table.expected {
			t.Errorf("Wrong url returned. Got %s, expected %s", result, table.expected)
		}
	}
}

func TestCreateRequest(t *testing.T) {
	ro := New("APIKEY12345", "api.rajaongkir.com/starter", nil)
	tables := []struct {
		method   string
		endpoint string
		payload  string
		isErr    bool
	}{
		{http.MethodGet, "/city", `{"foo":"bar"}`, false},
		{http.MethodPost, "/city", `{"foo":"bar"}`, false},
		{http.MethodPut, "/city", `{"foo":"bar"}`, false},
		{http.MethodDelete, "/city^", `{"foo":"bar"}`, false},
		{"øP", "/city", `{"foo":"bar"}`, true},
	}

	for _, table := range tables {
		req, err := ro.createRequest(table.method, table.endpoint, table.payload)

		isErr := false
		if err != nil {
			isErr = true
		}
		expectedIsErr := table.isErr
		if isErr != expectedIsErr {
			t.Errorf("Error mismatch. Got %s, expected %v", err, expectedIsErr)
		}

		if err == nil {
			method := req.Method
			expectedMethod := table.method
			if method != expectedMethod {
				t.Errorf("Method mismatch. Got %s, expected %s", method, expectedMethod)
			}

			key := req.Header.Get("key")
			expectedKey := "APIKEY12345"
			if key != expectedKey {
				t.Errorf("APIKey mismatch. Got %s, expected %s", key, expectedKey)
			}
		}

	}
}

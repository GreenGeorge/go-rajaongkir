package rajaongkir

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type received struct {
	receivedMethod   string
	receivedAPIKey   string
	receivedEndpoint string
}

func setupTest() (*httptest.Server, *RajaOngkir, *received) {
	rec := &received{}
	handler := func(w http.ResponseWriter, r *http.Request) {
		rec.receivedMethod = r.Method
		rec.receivedAPIKey = r.Header.Get("key")
		rec.receivedEndpoint = r.URL.String()
		fmt.Fprint(w, `{"rajaongkir": {"results": {}}}`)
	}
	ts := httptest.NewTLSServer(http.HandlerFunc(handler))
	testClient := ts.Client()
	hostname := strings.Replace(ts.URL, "https://", "", 1)
	ro := New("APIKEY12345", hostname, testClient)
	return ts, ro, rec
}

func TestNew(t *testing.T) {
	testClient := &http.Client{Timeout: time.Second * 5}
	ro := New("APIKEY12345", "test.com", testClient)

	expectedAPIKey := "APIKEY12345"
	expectedBaseURL := "test.com"
	expectedClientTimeout := time.Second * 5

	if ro.apiKey != expectedAPIKey {
		t.Errorf("Wrong API Key. Got %s, expected %s", ro.apiKey, expectedAPIKey)
	}
	if ro.baseURL != expectedBaseURL {
		t.Errorf("Wrong base URL. Got %s, expected %s", ro.baseURL, expectedBaseURL)
	}
	if ro.client.Timeout != expectedClientTimeout {
		t.Errorf("Wrong client timeout. Got %s, expected %s", ro.client.Timeout, expectedClientTimeout)
	}
}

func TestGetProvinces(t *testing.T) {
	ts, ro, rec := setupTest()
	defer ts.Close()
	ro.GetProvinces()
	expectedMethod := "GET"
	expectedAPIKey := "APIKEY12345"
	expectedEndpoint := "/province"

	if rec.receivedMethod != expectedMethod {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedMethod, expectedMethod)
	}
	if rec.receivedAPIKey != expectedAPIKey {
		t.Errorf("Wrong APIKEY. Received %s, expected %s", rec.receivedAPIKey, expectedAPIKey)
	}
	if rec.receivedEndpoint != expectedEndpoint {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedEndpoint, expectedEndpoint)
	}
}

func TestGetProvince(t *testing.T) {
	ts, ro, rec := setupTest()
	defer ts.Close()
	ro.GetProvince("41")
	expectedMethod := "GET"
	expectedAPIKey := "APIKEY12345"
	expectedEndpoint := "/province?id=41"

	if rec.receivedMethod != expectedMethod {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedMethod, expectedMethod)
	}
	if rec.receivedAPIKey != expectedAPIKey {
		t.Errorf("Wrong APIKEY. Received %s, expected %s", rec.receivedAPIKey, expectedAPIKey)
	}
	if rec.receivedEndpoint != expectedEndpoint {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedEndpoint, expectedEndpoint)
	}
}

func TestGetCities(t *testing.T) {
	ts, ro, rec := setupTest()
	defer ts.Close()
	ro.GetCities()
	expectedMethod := "GET"
	expectedAPIKey := "APIKEY12345"
	expectedEndpoint := "/city"

	if rec.receivedMethod != expectedMethod {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedMethod, expectedMethod)
	}
	if rec.receivedAPIKey != expectedAPIKey {
		t.Errorf("Wrong APIKEY. Received %s, expected %s", rec.receivedAPIKey, expectedAPIKey)
	}
	if rec.receivedEndpoint != expectedEndpoint {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedEndpoint, expectedEndpoint)
	}
}

func TestGetCost(t *testing.T) {
	ts, ro, rec := setupTest()
	defer ts.Close()
	ro.GetCost("153", "151", 1200, "jne")
	expectedMethod := "POST"
	expectedAPIKey := "APIKEY12345"
	expectedEndpoint := "/cost"

	if rec.receivedMethod != expectedMethod {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedMethod, expectedMethod)
	}
	if rec.receivedAPIKey != expectedAPIKey {
		t.Errorf("Wrong APIKEY. Received %s, expected %s", rec.receivedAPIKey, expectedAPIKey)
	}
	if rec.receivedEndpoint != expectedEndpoint {
		t.Errorf("Wrong method. Received %s, expected %s", rec.receivedEndpoint, expectedEndpoint)
	}
}

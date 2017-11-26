package rajaongkir

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Creates target URL for making the requests
func (r *RajaOngkir) createTargetURL(endpoint string) string {
	targetURL := fmt.Sprintf("https://%s%s", r.baseURL, endpoint)
	return targetURL
}

func (r *RajaOngkir) createRequest(method, endpoint string, payloadString string) (*http.Request, error) {
	url := r.createTargetURL(endpoint)
	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error in request", err)
	}
	req.Header.Set("key", r.apiKey)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	return req, err
}

func (r *RajaOngkir) executeRequest(req *http.Request) ([]byte, error) {
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

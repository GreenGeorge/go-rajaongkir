package rajaongkir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// List of endpoints according to https://rajaongkir.com/dokumentasi/starter
const (
	provinceEndpoint = "/province"
	cityEndpoint     = "/city"
	costEndpoint     = "/cost"
)

// RajaOngkir wraps our request operations
type RajaOngkir struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// New Creates a new Raja Ongkir API object
// containing the config
func New(apiKey, baseURL string, client *http.Client) *RajaOngkir {
	if client == nil {
		client = &http.Client{Timeout: time.Second * 10}
	}
	r := &RajaOngkir{apiKey, baseURL, client}
	return r
}

// Creates target URL for making the requests
func (r *RajaOngkir) createTargetURL(endpoint string) string {
	targetURL := fmt.Sprintf("https://%s%s", r.baseURL, endpoint)
	return targetURL
}

func (r *RajaOngkir) createRequest(method, endpoint string, payloadString string) *http.Request {
	url := r.createTargetURL(endpoint)
	payload := strings.NewReader(payloadString)
	req, reqErr := http.NewRequest(method, url, payload)
	if reqErr != nil {
		fmt.Println("Error in request", reqErr)
	}
	req.Header.Set("key", r.apiKey)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	return req
}

func (r *RajaOngkir) executeRequest(req *http.Request) []byte {
	res, resErr := r.client.Do(req)
	if resErr != nil {
		fmt.Println("Request failed", resErr)
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Error reading body", readErr)
	}
	return body
}

// GetProvinces fetches the list of provinces
func (r *RajaOngkir) GetProvinces() []map[string]string {
	req := r.createRequest(http.MethodGet, provinceEndpoint, "")
	body := r.executeRequest(req)
	var b map[string]map[string][]map[string]string
	json.Unmarshal(body, &b)
	provinces := b["rajaongkir"]["results"]
	return provinces
}

// GetProvince fetches a specific province
// given an ID
func (r *RajaOngkir) GetProvince(id string) map[string]interface{} {
	req := r.createRequest(http.MethodGet, fmt.Sprintf("%s?id=%s", provinceEndpoint, id), "")
	body := r.executeRequest(req)
	var b map[string]map[string]map[string]interface{}
	json.Unmarshal(body, &b)
	province := b["rajaongkir"]["results"]
	return province
}

// GetCities fetches the list of cities
func (r *RajaOngkir) GetCities() []map[string]string {
	req := r.createRequest(http.MethodGet, cityEndpoint, "")
	body := r.executeRequest(req)
	var b map[string]map[string][]map[string]string
	json.Unmarshal(body, &b)
	cities := b["rajaongkir"]["results"]
	return cities
}

// GetCost fetches the shipping rate
func (r *RajaOngkir) GetCost(origin, destination string, weight int, courier string) []map[string]string {
	queryString := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s", origin, destination, weight, courier)
	req := r.createRequest(http.MethodPost, costEndpoint, queryString)
	body := r.executeRequest(req)
	var b map[string]map[string][]map[string][]map[string]string
	json.Unmarshal(body, &b)
	// Access safely
	if len(b["rajaongkir"]["results"]) > 0 {
		results := b["rajaongkir"]["results"][0]
		costs := results["costs"]
		return costs
	}
	return []map[string]string{}
}

package rajaongkir

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// List of endpoints according to https://rajaongkir.com/dokumentasi/starter
const (
	provinceEndpoint = "/province"
	cityEndpoint     = "/city"
	costEndpoint     = "/cost"
)

type CitiesResp struct {
	RajaOngkir struct {
		Status struct {
			Code int `json:"code"`
		} `json:"status"`
		Results []City `json:"results"`
	} `json:"rajaongkir"`
}

type City struct {
	City       string `json:"city_name"`
	CityID     string `json:"city_id"`
	Province   string `json:"province"`
	ProvinceID string `json:"province_id"`
	Type       string `json:"type"`
}

type CostResp struct {
	RajaOngkir struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []Result `json:"results"`
	} `json:"rajaongkir"`
}

type Result struct {
	Code  string        `json:"code"`
	Name  string        `json:"name"`
	Costs []ServiceCost `json:"costs"`
}

type ServiceCost struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Cost struct {
	Value int64  `json:"value"`
	ETD   string `json:"etd"`
	Note  string `json:"note"`
}

// RajaOngkir struct wraps our request operations
type RajaOngkir struct {
	apiKey  string
	baseURL string
	request *gorequest.SuperAgent
}

// New Creates a new Raja Ongkir API object
// containing the config
func New(apiKey, baseURL string) *RajaOngkir {
	r := &RajaOngkir{apiKey, baseURL, gorequest.New()}
	return r
}

// Creates target URL for making the requests
func (r *RajaOngkir) createTargetURL(endpoint string) string {
	targetURL := fmt.Sprintf("https://%s%s", r.baseURL, endpoint)
	return targetURL
}

// Creates a get request with the proper headers
func (r *RajaOngkir) createGetRequest(endpoint string) *gorequest.SuperAgent {
	targetURL := r.createTargetURL(endpoint)
	return r.request.
		Get(targetURL).
		Set("key", r.apiKey)
}

// Creates a post request with the proper headers
// and the data loaded
func (r *RajaOngkir) createPostRequest(endpoint, data string) *gorequest.SuperAgent {
	targetURL := r.createTargetURL(endpoint)
	return r.request.
		Post(targetURL).
		Set("key", r.apiKey).
		Send(data)
}

// GetProvinces fetches the list of provinces
func (r *RajaOngkir) GetProvinces() string {
	request := r.createGetRequest(provinceEndpoint)
	_, body, err := request.End()
	if err != nil {
		fmt.Println("Request failed", err)
	}
	return body
}

// GetProvince fetches a specific province
// given an ID
func (r *RajaOngkir) GetProvince(id string) string {
	request := r.createGetRequest(fmt.Sprintf("%s?id=%s", provinceEndpoint, id))
	_, body, err := request.End()
	if err != nil {
		fmt.Println("Request failed", err)
	}
	return body
}

// GetCities fetches the list of cities
func (r *RajaOngkir) GetCities() CitiesResp {
	var b CitiesResp
	request := r.createGetRequest(cityEndpoint)
	_, _, err := request.EndStruct(&b)
	if err != nil {
		fmt.Println("Request failed", err)
	}
	return b
}

// GetCost fetches the shipping rate
func (r *RajaOngkir) GetCost(origin string, destination string, weight int, courier string) CostResp {
	var b CostResp
	data := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s", origin, destination, weight, courier)
	request := r.createPostRequest(costEndpoint, data)
	_, _, err := request.EndStruct(&b)
	if err != nil {
		fmt.Println("Request failed", err)
	}
	return b
}

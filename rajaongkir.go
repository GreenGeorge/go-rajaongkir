// Package rajaongkir provides methods for making requests to the RajaOngkir starter API.
// See https://rajaongkir.com/dokumentasi/starter for further details on the API and to get an API Key
package rajaongkir

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// List of endpoints according to https://rajaongkir.com/dokumentasi/starter
const (
	provinceEndpoint     = "/province"
	cityEndpoint         = "/city"
	costEndpoint         = "/cost"
	defaultClientTimeout = time.Second * 10
)

// RajaOngkir stores the credentials for accessing the API
type RajaOngkir struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

type query map[string]interface{}

type status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type carrierService struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []Cost `json:"costs"`
}

// Cost stores the details of the shipping cost
type Cost struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []struct {
		Value int    `json:"value"`
		ETD   string `json:"etd"`
		Note  string `json:"note"`
	} `json:"cost"`
}

// Province stores the details of a province
type Province struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

// City stores the details of a city
type City struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"city_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type provinceResponse struct {
	Rajaongkir struct {
		Query   query    `json:"query"`
		Status  status   `json:"status"`
		Results Province `json:"results"`
	} `json:"rajaongkir"`
}

type provincesResponse struct {
	Rajaongkir struct {
		Query   query      `json:"query"`
		Status  status     `json:"status"`
		Results []Province `json:"results"`
	} `json:"rajaongkir"`
}

type cityResponse struct {
	Rajaongkir struct {
		Query   query  `json:"query"`
		Status  status `json:"status"`
		Results City   `json:"results"`
	} `json:"rajaongkir"`
}

type citiesResponse struct {
	Rajaongkir struct {
		Query   query  `json:"query"`
		Status  status `json:"status"`
		Results []City `json:"results"`
	} `json:"rajaongkir"`
}

type costResponse struct {
	Rajaongkir struct {
		Query              query            `json:"query"`
		Status             status           `json:"status"`
		OriginDetails      City             `json:"origin_details"`
		DestinationDetails City             `json:"destination_details"`
		Results            []carrierService `json:"results"`
	} `json:"rajaongkir"`
}

// New initializes a new RajaOngkir struct
// with a default client configured if none is specified
func New(apiKey, baseURL string, client *http.Client) *RajaOngkir {
	if client == nil {
		client = &http.Client{Timeout: defaultClientTimeout}
	}
	r := &RajaOngkir{apiKey, baseURL, client}
	return r
}

func checkStatus(status *status) error {
	if status.Code >= 200 && status.Code < 300 {
		return nil
	}
	return errors.New(status.Description)
}

// GetProvinces fetches the list of provinces
func (r *RajaOngkir) GetProvinces() ([]Province, error) {
	re := &provincesResponse{}
	err := r.sendRequest(http.MethodGet, provinceEndpoint, "", re)
	if err != nil {
		return nil, err
	}
	err = checkStatus(&re.Rajaongkir.Status)
	if err != nil {
		return nil, err
	}
	provinces := re.Rajaongkir.Results
	return provinces, nil
}

// GetProvince fetches a specific province
// matching a given ID
func (r *RajaOngkir) GetProvince(id string) (Province, error) {
	re := &provinceResponse{}
	endpoint := fmt.Sprintf("%s?id=%s", provinceEndpoint, id)
	err := r.sendRequest(http.MethodGet, endpoint, "", re)
	if err != nil {
		return Province{}, err
	}
	err = checkStatus(&re.Rajaongkir.Status)
	if err != nil {
		return Province{}, err
	}
	province := re.Rajaongkir.Results
	return province, nil
}

// GetCities fetches the list of cities
func (r *RajaOngkir) GetCities() ([]City, error) {
	re := &citiesResponse{}
	err := r.sendRequest(http.MethodGet, cityEndpoint, "", re)
	if err != nil {
		return nil, err
	}
	cities := re.Rajaongkir.Results
	return cities, nil
}

// GetCost fetches the shipping rate
// given the origin, destination, weight, and courier service
func (r *RajaOngkir) GetCost(origin, destination string, weight int, courier string) ([]Cost, error) {
	queryString := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s", origin, destination, weight, courier)
	re := &costResponse{}
	err := r.sendRequest(http.MethodPost, costEndpoint, queryString, re)
	if err != nil {
		return nil, err
	}
	err = checkStatus(&re.Rajaongkir.Status)
	if err != nil {
		return nil, err
	}
	costs := re.Rajaongkir.Results[0].Costs
	return costs, nil
}

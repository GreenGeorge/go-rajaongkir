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

const provinceRes string = `{
    "rajaongkir": {
        "query": {
            "id": "12"
        },
        "status": {
            "code": 200,
            "description": "OK"
        },
        "results": {
           "province_id": "12",
           "province": "Kalimantan Barat"
        }
    }
}`

const provincesRes string = `{
    "rajaongkir": {
        "query": {
            "id": "12"
        },
        "status": {
            "code": 400,
            "description": "OK"
        },
        "results": [{
           "province_id": "12",
           "province": "Kalimantan Barat"
        },{
						"province_id": "13",
						"province": "Kalimantan Timur"
				}]
    }
}`

const cityRes string = `{
    "rajaongkir": {
        "query": {
            "province": "5",
            "id": "39"
        },
        "status": {
            "code": 200,
            "description": "OK"
        },
        "results": {
           "city_id": "39",
           "province_id": "5",
           "province": "DI Yogyakarta",
           "type": "Kabupaten",
           "city_name": "Bantul",
           "postal_code": "55700"
        }
    }
}`

const citiesRes string = `{
    "rajaongkir": {
        "query": {
            "province": "5",
            "id": "39"
        },
        "status": {
            "code": 200,
            "description": "OK"
        },
        "results": [{
           "city_id": "39",
           "province_id": "5",
           "province": "DI Yogyakarta",
           "type": "Kabupaten",
           "city_name": "Bantul",
           "postal_code": "55700"
        }]
    }
}`

const costRes string = `
{
   "rajaongkir":{
      "query":{
         "origin":"501",
         "destination":"114",
         "weight":1700,
         "courier":"jne"
      },
      "status":{
         "code":200,
         "description":"OK"
      },
      "origin_details":{
         "city_id":"501",
         "province_id":"5",
         "province":"DI Yogyakarta",
         "type":"Kota",
         "city_name":"Yogyakarta",
         "postal_code":"55000"
      },
      "destination_details":{
         "city_id":"114",
         "province_id":"1",
         "province":"Bali",
         "type":"Kota",
         "city_name":"Denpasar",
         "postal_code":"80000"
      },
      "results":[
         {
            "code":"jne",
            "name":"Jalur Nugraha Ekakurir (JNE)",
            "costs":[
               {
                  "service":"OKE",
                  "description":"Ongkos Kirim Ekonomis",
                  "cost":[
                     {
                        "value":38000,
                        "etd":"4-5",
                        "note":""
                     }
                  ]
               },
               {
                  "service":"REG",
                  "description":"Layanan Reguler",
                  "cost":[
                     {
                        "value":44000,
                        "etd":"2-3",
                        "note":""
                     }
                  ]
               },
               {
                  "service":"SPS",
                  "description":"Super Speed",
                  "cost":[
                     {
                        "value":349000,
                        "etd":"",
                        "note":""
                     }
                  ]
               },
               {
                  "service":"YES",
                  "description":"Yakin Esok Sampai",
                  "cost":[
                     {
                        "value":98000,
                        "etd":"1-1",
                        "note":""
                     }
                  ]
               }
            ]
         }
      ]
   }
}`

func setupTest(jsonResponse string) (*httptest.Server, *RajaOngkir, *received) {
	rec := &received{}
	handler := func(w http.ResponseWriter, r *http.Request) {
		rec.receivedMethod = r.Method
		rec.receivedAPIKey = r.Header.Get("key")
		rec.receivedEndpoint = r.URL.String()
		fmt.Fprint(w, jsonResponse)
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
	rodc := New("APIKEY12345", "test.com", nil)

	expectedAPIKey := "APIKEY12345"
	expectedBaseURL := "test.com"
	expectedClientTimeout := time.Second * 5
	expectedDefaultClientTimeout := time.Second * 10

	if ro.apiKey != expectedAPIKey {
		t.Errorf("Wrong API Key. Got %s, expected %s", ro.apiKey, expectedAPIKey)
	}
	if ro.baseURL != expectedBaseURL {
		t.Errorf("Wrong base URL. Got %s, expected %s", ro.baseURL, expectedBaseURL)
	}
	if ro.client.Timeout != expectedClientTimeout {
		t.Errorf("Wrong client timeout. Got %s, expected %s", ro.client.Timeout, expectedClientTimeout)
	}
	if rodc.client.Timeout != expectedDefaultClientTimeout {
		t.Errorf("Default client not set. Got %s, expected %s timeout", rodc.client.Timeout, expectedDefaultClientTimeout)
	}
}

func TestGetProvinces(t *testing.T) {
	ts, ro, rec := setupTest(provincesRes)
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
	ts, ro, rec := setupTest(provinceRes)
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
	ts, ro, rec := setupTest(citiesRes)
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
	ts, ro, rec := setupTest(costRes)
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

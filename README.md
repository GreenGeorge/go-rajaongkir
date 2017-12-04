# Go RajaOngkir

Simple way to make requests to the [**RajaOngkir API**][1]. Uses Go's `net/http`. Inspired by [`rapito/go-shopify`][3]

[![CircleCI](https://circleci.com/gh/GreenGeorge/go-rajaongkir.svg?style=shield)](https://circleci.com/gh/GreenGeorge/go-rajaongkir)
[![codecov](https://codecov.io/gh/GreenGeorge/go-rajaongkir/branch/master/graph/badge.svg)](https://codecov.io/gh/GreenGeorge/go-rajaongkir)
[![GoDoc](https://godoc.org/github.com/GreenGeorge/go-rajaongkir?status.svg)](http://godoc.org/github.com/GreenGeorge/go-rajaongkir)
[![goreport](https://goreportcard.com/badge/github.com/GreenGeorge/go-rajaongkir)](https://goreportcard.com/report/github.com/GreenGeorge/go-rajaongkir)


[1]: https://rajaongkir.com/dokumentasi
[2]: https://github.com/parnurzeal/gorequest
[3]: https://github.com/rapito/go-shopify

## Installation
```
go get github.com/GreenGeorge/go-rajaongkir
```

## Usage
```go
  ...
  import "github.com/GreenGeorge/go-rajaongkir"
  ...

  const apiKey = "YOUR_API_KEY_HERE"
  const baseURL = "api.rajaongkir.com/starter"

  // Initialize RajaOngkir
  // BYO http.Client if you wish. Pass it as the 3rd parameter
  // otherwise go-rajaongkir will preconfigure one for you
  r := rajaongkir.New(apiKey, baseURL, nil)

  // Get a list of provinces
  // Returns []Province
  provinces, err := r.GetProvinces()

  // Get a list of cities
  // Returns []City
  cities, err := r.GetCities()

  origin      := 501      // origin province code
  destination := 114      // destination province code
  weight      := 1700     // weight in grams
  courier     := "jne"    // delivery service

  // Get the shipping cost
  // Right now you can only pass JNE as courier
  // Returns []Cost
  shippingCosts, err := r.GetCost(origin, destination, weight, courier)
  ...
```

## Contributing
Got ideas? Open an issue for discussion. Contributions are always welcome. Send a PR with tests.

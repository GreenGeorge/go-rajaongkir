# Go RajaOngkir

Simple way to make requests to the [**RajaOngkir API**][1]. Uses [`parnurzeal/gorequest`][2] and inspired by [`rapito/go-shopify`][3]

[1]: https://rajaongkir.com/dokumentasi
[2]: https://github.com/parnurzeal/gorequest
[3]: https://github.com/rapito/go-shopify

## Installation
```
go get github.com/GreenGeorge/go-rajaongkir
```

## Usage
```
  ...
  import "github.com/GreenGeorge/rajaongkir"
  ...

  const apiKey = "YOUR_API_KEY_HERE"
  const baseURL = "api.rajaongkir.com/starter"

  r := rajaongkir.New(apiKey, baseURL)

  // Get a list of provinces
  provinces := r.GetProvinces()

  // Get a list of cities
  cities := r.GetCities()

  origin      := 501      // origin province code
  destination := 114      // destination province code
  weight      := 1700     // weight in grams
  courier     := "jne"    // delivery service

  // Get the shipping cost
  shippingCost := r.GetCost(origin, destination, weight, courier)
  ...
```

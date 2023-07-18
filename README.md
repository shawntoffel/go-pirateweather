# go-pirateweather
[![Go Reference](https://pkg.go.dev/badge/github.com/shawntoffel/go-pirateweather.svg)](https://pkg.go.dev/github.com/shawntoffel/go-pirateweather) 
 [![Go Report Card](https://goreportcard.com/badge/github.com/shawntoffel/go-pirateweather)](https://goreportcard.com/report/github.com/shawntoffel/go-pirateweather) [![Build status](https://github.com/shawntoffel/go-pirateweather/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/shawntoffel/go-pirateweather/actions/workflows/go.yml)
 
A [Pirate Weather](https://pirateweather.net) API client in Go.

<img align="right" alt="go-pirateweather logo" src="https://user-images.githubusercontent.com/2343437/214456395-74a98129-32dd-4cc5-b83f-e1e8ae89d9c7.png" width="350">

 ## Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/shawntoffel/go-pirateweather
```

## Usage

Import the package into your project:

```go
import (
    "github.com/shawntoffel/go-pirateweather"
)
```

Create a new pirateweather client:

```go
client := pirateweather.Client{}
```

Build a request:

```go
request := pirateweather.ForecastRequest{
  Latitude:  40.7128,
  Longitude: -74.0059,
}
```

Get a forecast:
```go
ctx := context.Background()

// The api token can be found on your Pirate Weather dashboard.
response, err := client.Forecast(ctx, apiToken, request)
```

## Documentation
- [![Go Reference](https://pkg.go.dev/badge/github.com/shawntoffel/go-pirateweather.svg)](https://pkg.go.dev/github.com/shawntoffel/go-pirateweather) 
- [Getting Started](https://pirateweather.net/en/latest/)
- [API](https://pirateweather.net/en/latest/API/)

## Examples
- [Simple forecast](https://github.com/shawntoffel/go-pirateweather/tree/master/examples/simple/main.go)
- [Additional Forecast Options](https://github.com/shawntoffel/go-pirateweather/tree/master/examples/forecast_options/main.go)
- [Custom HTTP Client](https://github.com/shawntoffel/go-pirateweather/tree/master/examples/custom_http_client/main.go)

## Troubleshooting
Please use the GitHub [Discussions](https://github.com/shawntoffel/go-pirateweather/discussions) tab for questions regarding this client library.

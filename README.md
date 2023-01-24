# go-pirateweather


A [Pirate Weather](https://pirateweather.net) API client in Go.

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
- [Getting Started](https://pirateweather.net/getting-started)
- [API](https://docs.pirateweather.net/en/latest/)

## Examples
- [Simple forecast](https://github.com/shawntoffel/go-pirateweather/tree/master/examples/simple/main.go)
- [Additional Forecast Options](https://github.com/shawntoffel/go-pirateweather/tree/master/examples/forecast_options/main.go)
- [Custom HTTP Client](https://github.com/shawntoffel/go-pirateweather/tree/master/examples/custom_http_client/main.go)

## Troubleshooting
Please use the GitHub [Discussions](https://github.com/shawntoffel/go-pirateweather/discussions) tab for questions regarding this client library.

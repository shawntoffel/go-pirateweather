package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shawntoffel/go-pirateweather"
)

func main() {
	apiToken := flag.String("t", "", "API Token")
	flag.Parse()

	ctx := context.Background()

	// Create a Pirate Weather client
	client := pirateweather.Client{

		// Provide a custom HTTP client
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	request := pirateweather.ForecastRequest{
		Latitude:  40.7128,
		Longitude: -74.0059,
	}

	response, err := client.Forecast(ctx, *apiToken, request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(response.Currently.Temperature)
}

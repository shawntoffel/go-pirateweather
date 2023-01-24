package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/shawntoffel/go-pirateweather"
)

func main() {
	apiToken := flag.String("t", "", "API Token")
	flag.Parse()

	ctx := context.Background()

	// Create a Pirate Weather client
	client := pirateweather.Client{}

	request := pirateweather.ForecastRequest{
		Latitude:  40.7128,
		Longitude: -74.0059,

		Options: pirateweather.ForecastRequestOptions{
			// Exclude hourly & minutely forecasts in the response
			Exclude: "hourly,minutely",

			// Return data in SI units.
			Units: "si",
		},
	}

	response, err := client.Forecast(ctx, *apiToken, request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(response.Currently.Temperature)
}

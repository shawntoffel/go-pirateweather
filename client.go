package pirateweather

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

var (
	// Pirate Weather API base URL
	BaseUrl = "https://api.pirateweather.net"

	// UserAgent to send along with requests.
	UserAgent = "shawntoffel/go-pirateweather"
)

// Client is a Pirate Weather API client.
type Client struct {
	HttpClient *http.Client
}

// Forecast retrieves a forecast.
func (d *Client) Forecast(ctx context.Context, token string, request ForecastRequest) (*ForecastResponse, error) {
	response := ForecastResponse{}
	err := d.get(ctx, token, request, &response)
	return &response, err
}

func (d *Client) get(ctx context.Context, token string, request urlBuilder, output interface{}) error {
	if d.HttpClient == nil {
		d.HttpClient = &http.Client{}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, request.url(token), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("User-Agent", UserAgent)

	response, err := d.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = validateResponse(response)
	if err != nil {
		return err
	}

	return decode(response, &output)
}

func validateResponse(response *http.Response) error {
	if response.StatusCode == http.StatusOK {
		return nil
	}

	errorResponse := ErrorResponse{}

	err := decode(response, &errorResponse)
	if err != nil {
		return err
	}

	return &RestError{
		Response:      response,
		ErrorResponse: &errorResponse,
	}
}

func decode(response *http.Response, into interface{}) error {
	body, err := decompress(response)
	if err != nil {
		return err
	}

	return unmarshal(body, into)
}

func decompress(response *http.Response) (io.Reader, error) {
	header := response.Header.Get("Content-Encoding")
	if len(header) < 1 {
		return response.Body, nil
	}

	return gzip.NewReader(response.Body)
}

func unmarshal(body io.Reader, into interface{}) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if bytes == nil || len(bytes) < 1 {
		return nil
	}

	return json.Unmarshal(bytes, &into)
}

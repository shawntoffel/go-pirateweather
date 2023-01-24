package pirateweather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestParseFullForecastResponse(t *testing.T) {
	forecast(t, "testdata/sanitized_full_weather.json")
}

func TestParseExcludedForecastResponse(t *testing.T) {
	forecast(t, "testdata/sanitized_all_excluded.json")
}

func TestErrorResponse(t *testing.T) {
	client := Client{}

	server, expected, err := getMockServerWithFileData("testdata/error.json", http.StatusNotFound)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	_, err = client.Forecast(context.TODO(), "", ForecastRequest{})
	if err == nil {
		t.Errorf("expected request to error")
		return
	}

	restError, ok := err.(*RestError)
	if !ok {
		t.Errorf("expected error to be of type %s", reflect.TypeOf(RestError{}))
		return
	}

	actual, err := json.Marshal(restError.ErrorResponse)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func forecast(t *testing.T, filename string) {
	client := Client{}

	server, expected, err := getMockServerWithFileData(filename, http.StatusOK)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	response, err := client.Forecast(context.TODO(), "", ForecastRequest{})
	if err != nil {
		t.Error(err.Error())
	}

	actual, err := json.Marshal(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func assertJsonEqual(t *testing.T, expected []byte, actual []byte) {
	var a, b interface{}

	err := json.Unmarshal(expected, &a)
	if err != nil {
		t.Error(err.Error())
	}

	err = json.Unmarshal(actual, &b)
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("\nexpected: %+v\nactual: %+v", a, b)
	}
}

func getMockServerWithFileData(filename string, statusCode int) (*httptest.Server, []byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, bytes, err
	}

	return getMockServer(bytes, statusCode), bytes, nil
}

func getMockServer(data []byte, statusCode int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, string(data))
	}))

	return server
}

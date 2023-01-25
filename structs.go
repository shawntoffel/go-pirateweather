package pirateweather

import (
	"fmt"
	"net/http"
	"time"
)

// ForecastRequest contains all available options for requesting a forecast
type ForecastRequest struct {
	Latitude  float64
	Longitude float64
	Time      time.Time
	Options   ForecastRequestOptions
}

// ForecastRequestOptions are optional and passed as query parameters
type ForecastRequestOptions struct {
	Exclude string
	Extend  string
	Lang    string
	Units   string
}

// ForecastResponse is the response containing all requested properties
type ForecastResponse struct {
	Latitude  float64    `json:"latitude,omitempty"`
	Longitude float64    `json:"longitude,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
	Offset    float64    `json:"offset,omitempty"`
	Elevation float64    `json:"elevation,omitempty"`
	Currently *DataPoint `json:"currently,omitempty"`
	Minutely  *DataBlock `json:"minutely,omitempty"`
	Hourly    *DataBlock `json:"hourly,omitempty"`
	Daily     *DataBlock `json:"daily,omitempty"`
	Alerts    []*Alert   `json:"alerts,omitempty"`
	Flags     *Flags     `json:"flags,omitempty"`
}

// DataPoint contains various properties, each representing the average (unless otherwise specified) of a particular weather phenomenon occurring during a period of time.
type DataPoint struct {
	ApparentTemperature         float64 `json:"apparentTemperature,omitempty"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime,omitempty"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime,omitempty"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	NearestStormBearing         float64 `json:"nearestStormBearing,omitempty"`
	NearestStormDistance        float64 `json:"nearestStormDistance,omitempty"`
	Ozone                       float64 `json:"ozone,omitempty"`
	PrecipAccumulation          float64 `json:"precipAccumulation,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityError        float64 `json:"precipIntensityError,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	SunriseTime                 int64   `json:"sunriseTime,omitempty"`
	SunsetTime                  int64   `json:"sunsetTime,omitempty"`
	Temperature                 float64 `json:"temperature,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         int64   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          int64   `json:"temperatureLowTime,omitempty"`
	TemperatureMax              float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime,omitempty"`
	TemperatureMin              float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime          int64   `json:"temperatureMinTime,omitempty"`
	Time                        int64   `json:"time,omitempty"`
	UvIndex                     float64 `json:"uvIndex,omitempty"`
	UvIndexTime                 int64   `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	WindBearing                 float64 `json:"windBearing,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                int64   `json:"windGustTime,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
}

// DataBlock represents the various weather phenomena occurring over a period of time
type DataBlock struct {
	Data    []DataPoint `json:"data,omitempty"`
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
}

// Alert contains objects representing the severe weather warnings issued for the requested location by a governmental authority
type Alert struct {
	Description string   `json:"description,omitempty"`
	Expires     int64    `json:"expires,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Time        int64    `json:"time,omitempty"`
	Title       string   `json:"title,omitempty"`
	Uri         string   `json:"uri,omitempty"`
}

// Flags contains various metadata information related to the request
type Flags struct {
	NearestStation float64           `json:"nearest-station,omitempty"`
	Sources        []string          `json:"sources,omitempty"`
	Units          string            `json:"units,omitempty"`
	SourceTimes    map[string]string `json:"sourceTimes,omitempty"`
	Version        string            `json:"version,omitempty"`
}

// ErrorResponse is returned in response to an API error.
type ErrorResponse struct {
	Message      string `json:"message,omitempty"`
	Subscription string `json:"Subscription,omitempty"`
	Help         string `json:"Help,omitempty"`
}

// RestError has both http response & deserialized content error details.
type RestError struct {
	Response      *http.Response
	ErrorResponse *ErrorResponse
}

func (e *RestError) Error() string {
	msg := fmt.Sprintf("pirateweather: status code: %d %s", e.Response.StatusCode, http.StatusText(e.Response.StatusCode))

	if e.ErrorResponse == nil {
		return msg
	}

	if len(e.ErrorResponse.Message) > 0 {
		msg += ", message: " + e.ErrorResponse.Message
	}

	if len(e.ErrorResponse.Subscription) > 0 {
		msg += ", subscription: " + e.ErrorResponse.Subscription
	}

	if len(e.ErrorResponse.Help) > 0 {
		msg += ", help: " + e.ErrorResponse.Help
	}

	return msg
}

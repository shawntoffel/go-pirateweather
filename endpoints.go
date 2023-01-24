package pirateweather

import (
	"fmt"
	"net/url"
)

// Construct the forecast URL
func (r ForecastRequest) url(token string) string {
	endpoint := BaseUrl + "/forecast/" + fmt.Sprintf("%s/%g,%g", token, r.Latitude, r.Longitude)

	if !r.Time.IsZero() {
		endpoint += fmt.Sprintf(",%d", r.Time.Unix())
	}

	return endpoint + encodeUrlParameters(r.Options.urlValues())
}

// Create optional url Values for encoding into query string parameters (exclude=hourly&units=si)
func (o ForecastRequestOptions) urlValues() url.Values {
	q := url.Values{}

	if o.Exclude != "" {
		q.Add("exclude", o.Exclude)
	}
	if o.Extend != "" {
		q.Add("extend", o.Extend)
	}
	if o.Lang != "" {
		q.Add("lang", o.Lang)
	}
	if o.Units != "" {
		q.Add("units", o.Units)
	}

	return q
}

func encodeUrlParameters(values url.Values) string {
	queryString := values.Encode()

	if queryString == "" {
		return queryString
	}

	return "?" + queryString
}

type urlBuilder interface {
	url(token string) string
}

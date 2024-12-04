package models

// return value from locations api request, add struct tag
type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// api return a json object contains a list of locations, we need this struct to help parse the response
type LocationResponse struct {
	Locations []Location `json:"locations"`
}

type ForecastProperties struct {
	ForecastURL string `json:"forecast"`
}

// represent the whole json object
//
//	eg: "properties": {
//	    "forecast": "https://api.weather.gov/gridpoints/PSR/166,60/forecast"
//	}
type PointsResponse struct {
	Properties ForecastProperties `json:"properties"`
}

type ForecastPeriod struct {
	DetailedForecast string `json:"detailedForecast"`
}

type ForecastPropertiesDetail struct {
	Periods []ForecastPeriod `json:"periods"`
}

type ForecastResponse struct {
	Properties ForecastPropertiesDetail `json:"properties"`
}

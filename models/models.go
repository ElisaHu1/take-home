package models

// return value from locations api request
type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type Weather struct {
	Forecast            string `json:"forecast"`
	IssuingOffice       string `json:"issuing_office"`
	ObservationStations string `json:"observation_stations"`
	CountyZone          string `json:"county_zone"`
	FireZone            string `json:"fire_zone"`
}

package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const nwsURL = "https://api.weather.gov/points/"

func fetchWeatherForecast(latitude float64, longitude float64) (string, error) {
	pointsURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude)
	// step 1 fetch weather forecast: curl -sL https://api.weather.gov/points/33.50921,-111.89903 | jq '.properties.forecast'
	resp, err := http.Get(pointsURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// log.Println("Raw JSON response:", string(body))

	var pointsResponse PointsResponse
	err = json.Unmarshal(body, &pointsResponse)
	if err != nil {
		return "", err
	}

	forecastURL := pointsResponse.Properties.ForecastURL

	// step 2 fetch weather forecast, curl -s https://api.weather.gov/gridpoints/PSR/166,60/forecast
	resp, err = http.Get(forecastURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var forecastResponse ForecastResponse
	err = json.Unmarshal(body, &forecastResponse)
	if err != nil {
		return "", err
	}

	// | jq '.properties.periods[0].detailedForecast'
	if len(forecastResponse.Properties.Periods) == 0 {
		return "", fmt.Errorf("no forecast periods found")
	}

	return forecastResponse.Properties.Periods[0].DetailedForecast, nil
}

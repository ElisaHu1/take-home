package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		location, err := fetchRandomLocation()
		if err != nil {
			http.Error(w, "Failed to fetch location", http.StatusInternalServerError)
			return
		}

		forecast, err := fetchWeatherForecast(location.Latitude, location.Longitude)
		if err != nil {
			http.Error(w, "Failed to fetch weather forecast", http.StatusInternalServerError)
			return
		}

		response := fmt.Sprintf("The weather in %s is: %s", location.Name, forecast)
		w.Write([]byte(response))
	})

	log.Println("Server started at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// func fetchRandomLocation() (Location, error) {
// 	resp, err := http.Get("https://locations.patch3s.dev/api/random")
// 	if err != nil {
// 		return Location{}, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return Location{}, err
// 	}

// 	var locationResponse LocationResponse
// 	err = json.Unmarshal(body, &locationResponse)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return Location{}, err
// 	}

// 	fmt.Println("after unmarshal:", locationResponse)

// 	if len(locationResponse.Locations) == 0 {
// 		return Location{}, fmt.Errorf("no locations found")
// 	}

// 	return locationResponse.Locations[0], nil
// }

// func fetchWeatherForecast(latitude float64, longitude float64) (string, error) {
// 	pointsURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude)
// 	// step 1 fetch weather forecast: curl -sL https://api.weather.gov/points/33.50921,-111.89903 | jq '.properties.forecast'
// 	resp, err := http.Get(pointsURL)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	// log.Println("Raw JSON response:", string(body))

// 	var pointsResponse PointsResponse
// 	err = json.Unmarshal(body, &pointsResponse)
// 	if err != nil {
// 		return "", err
// 	}

// 	forecastURL := pointsResponse.Properties.ForecastURL

// 	// step 2 fetch weather forecast, curl -s https://api.weather.gov/gridpoints/PSR/166,60/forecast
// 	resp, err = http.Get(forecastURL)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err = io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	var forecastResponse ForecastResponse
// 	err = json.Unmarshal(body, &forecastResponse)
// 	if err != nil {
// 		return "", err
// 	}

// 	// | jq '.properties.periods[0].detailedForecast'
// 	if len(forecastResponse.Properties.Periods) == 0 {
// 		return "", fmt.Errorf("no forecast periods found")
// 	}

// 	return forecastResponse.Properties.Periods[0].DetailedForecast, nil
// }

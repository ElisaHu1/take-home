package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elisahu1/take-home/models"
)

const nwsURL = "https://api.weather.gov/points/"

func FetchWeather(latitude float64, longitude float64) (models.Weather, error) {
	weatherURL := fmt.Sprintf("%s%.4f,%.4f", nwsURL, latitude, longitude)
	response, err := http.Get(weatherURL)
	if err != nil {
		return models.Weather{}, fmt.Errorf("Get failed: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Weather{}, fmt.Errorf("NO 200: %v", err)
	}

	var pointsData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&pointsData)
	if err != nil {
		return models.Weather{}, fmt.Errorf("failed to decode points response: %v", err)
	}

	// Log the response
	responseJSON, err := json.MarshalIndent(pointsData, "", "  ") // Indent for readability
	if err != nil {
		log.Printf("Failed to marshal response for logging: %v", err)
	} else {
		log.Printf("API Response:\n%s", string(responseJSON))
	}

	return models.Weather{}, nil

}

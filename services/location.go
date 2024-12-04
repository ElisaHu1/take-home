package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

const locationURL = "https://locations.patch3s.dev/api/"

func GetRandomLocation() (models.Location, error) {
	resp, err := http.Get(locationURL)
	if err != nil {
		return models.Location{}, fmt.Errorf("no response: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Location{}, fmt.Errorf("error code: %v", err)
	}

	var location models.Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return models.Location{}, fmt.Errorf("fail to decode: %v", err)
	}

	return location, nil
}

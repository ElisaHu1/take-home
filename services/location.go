package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/elisahu1/take-home/models"
)

const locationURL = "https://locations.patch3s.dev/api/"

func FetchRandomLocation() (models.Location, error) {
	resp, err := http.Get(locationURL + "random")
	if err != nil {
		return models.Location{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Location{}, err
	}

	// fmt.Println("raw Body:", string(body))
	// raw body example: {"locations":[{"name":"Caldwell","latitude":40.83982,"longitude":-74.27654}]}

	var locationResponse models.LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		fmt.Println("err:", err)
		return models.Location{}, err
	}

	// fmt.Println("after unmarshal:", locationResponse)
	// {[{Caldwell 40.83982 -74.27654}]}

	if len(locationResponse.Locations) == 0 {
		return models.Location{}, fmt.Errorf("no locations found")
	}

	return locationResponse.Locations[0], nil
}

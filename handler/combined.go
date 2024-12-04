package handler

import (
	"fmt"
	"net/http"

	"../services"
)

func CombinedHandler(w http.ResponseWriter, r *http.Request) {
	location, err := services.GetRandomLocation()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching location: %v", err), http.StatusInternalServerError)
		return
	}

	weather, err := services.FetchWeather(location.Latitude, location.Longitude)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching weather: %v", err), http.StatusInternalServerError)
		return
	}

	// result := models.CombinedResult{
	// 	Location: location,
	// 	Weather:  weather,
	// }

	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(result); err != nil {
	// 	http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	// }
}

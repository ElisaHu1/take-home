package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elisahu1/take-home/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mocking FetchRandomLocation
func MockFetchRandomLocation() (*models.Location, error) {
	return &models.Location{
		Name:      "Antigo",
		Latitude:  45.1407,
		Longitude: -88.4343,
	}, nil
}

// Mocking FetchWeatherForecast
func MockFetchWeatherForecast(latitude, longitude float64) (string, error) {
	return "Snow likely. Mostly cloudy. High near 28, with temperatures falling to around 16 in the afternoon. West northwest wind 17 to 24 mph, with gusts as high as 38 mph. Chance of precipitation is 50%. New snow accumulation of less than one inch possible.", nil
}

func TestWeatherHandler(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Initialize Gin router
	r := gin.Default()

	// Set up routes (use mocked functions here)
	r.GET("/", func(c *gin.Context) {
		location, err := MockFetchRandomLocation()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to get random location"})
			return
		}

		forecast, err := MockFetchWeatherForecast(location.Latitude, location.Longitude)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to get weather report"})
			return
		}

		response := fmt.Sprintf("The weather in %s is: %s", location.Name, forecast)
		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
	})

	// Create a test HTTP request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve HTTP request
	r.ServeHTTP(rr, req)

	// Assert Status Code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert Response Body
	expectedMessage := "The weather in Antigo is: Snow likely. Mostly cloudy. High near 28, with temperatures falling to around 16 in the afternoon. West northwest wind 17 to 24 mph, with gusts as high as 38 mph. Chance of precipitation is 50%. New snow accumulation of less than one inch possible."
	assert.Contains(t, rr.Body.String(), expectedMessage)

	// Optionally, assert that the JSON response contains the "message" key
	var responseMap map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		t.Fatalf("could not parse response body: %v", err)
	}
	assert.Contains(t, responseMap, "message")
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elisahu1/take-home/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

var log zerolog.Logger

func init() {
	// Set up structured logging to the console
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func main() {

	var limiter = rate.NewLimiter(rate.Limit(10), 1) // 10 requests per second, with a burst size of 1
	r := gin.Default()

	// in production, we do not want to trust *,
	// r.SetTrustedProxies([]string{
	// 	"10.0.0.1", // ip range for reserve proxy, ngnix
	// 	"192.168.1.0/24", // ip range for internal proxies
	// })

	r.Use(func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "too much request, try again later"})
			ctx.Abort()
			return
		}
		ctx.Next()
	})

	r.GET("/", func(c *gin.Context) {
		location, err := services.FetchRandomLocation()
		if err != nil {
			log.Error().Err(err).Msg("Failed to fetch random location")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to get random location"})
			return
		}

		forecast, err := services.FetchWeatherForecast(location.Latitude, location.Longitude)
		if err != nil {
			log.Error().Err(err).Msg("Failed to fetch weather forecast")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to get weather report"})
			return
		}

		response := fmt.Sprintf("The weather in %s is:%s", location.Name, forecast)

		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
	})

	log.Println("Server started at :5000")
	err := r.Run(":5000")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
		os.Exit(1)
	}
}

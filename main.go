package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/mic90/metrics-api/api/http"
	"github.com/mic90/metrics-api/api/http/health"
	"github.com/mic90/metrics-api/api/http/metrics"
	_ "github.com/mic90/metrics-api/docs"
	"github.com/mic90/metrics-api/persistance/driver"
	"log"
)

// Setup configures HTTP api endpoints
// @title Metrics API
// @version 1.0
// @description metrics api supports storage and retrieval of various time-series metrics
// @host localhost:8080
// @BasePath /api/v1/
func main() {
	// setup health routes
	healthService := &health.Service{}
	// setup metrics routes
	storage := driver.NewMemory()
	metricsService := metrics.NewService(storage)

	app := http.Setup(healthService, metricsService)

	app.Get("/swagger/*", swagger.Handler)

	log.Fatal(app.Listen(":8080"))
}

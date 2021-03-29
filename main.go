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

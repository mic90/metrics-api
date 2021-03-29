package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/mic90/metrics-api/api/http/health"
	"github.com/mic90/metrics-api/api/http/metrics"
	_ "github.com/mic90/metrics-api/docs"
	"github.com/mic90/metrics-api/persistance/driver"
	"log"
)

// @title Metrics API
// @version 1.0
// @description metrics api supports storage and retrieval of various time-series metrics
// @host localhost:8080
// @BasePath /api/v1/
func main() {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	api := app.Group("api").Group("v1")

	app.Get("/swagger/*", swagger.Handler)

	// setup health routes
	healthService := health.Service{}
	health.Router(api, healthService)

	// setup metrics routes
	storage := driver.NewMemory()
	metricsService := metrics.NewService(storage)
	metrics.Router(api, metricsService)

	log.Fatal(app.Listen(":8080"))
}

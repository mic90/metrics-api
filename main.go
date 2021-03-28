package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/mic90/metrics-api/api/http/router"
	"github.com/mic90/metrics-api/api/http/services"
	"github.com/mic90/metrics-api/persistance/driver"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	api := app.Group("api").Group("v1")

	// setup health routes
	healthService := services.HealthService{}
	router.HealthRouter(api, healthService)

	// setup metrics routes
	storage := driver.NewMemory()
	metricsService := services.NewMetricService(storage)
	router.MetricRouter(api, metricsService)

	log.Fatal(app.Listen(":8080"))
}

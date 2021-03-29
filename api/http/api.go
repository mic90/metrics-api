package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/mic90/metrics-api/api/http/health"
	"github.com/mic90/metrics-api/api/http/metrics"
)

// Setup configures HTTP api endpoints
// @title Metrics API
// @version 1.0
// @description metrics api supports storage and retrieval of various time-series metrics
// @host localhost:8080
// @BasePath /api/v1/
func Setup(healthSvc *health.Service, metricSvc *metrics.Service) *fiber.App {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	api := app.Group("api").Group("v1")

	// setup health routes
	health.Router(api, healthSvc)

	// setup metrics routes
	metrics.Router(api, metricSvc)

	return app
}
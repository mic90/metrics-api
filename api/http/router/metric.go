package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mic90/metrics-api/api/http/services"
)

// MetricRouter contains all the metrics routes
func MetricRouter(app fiber.Router, svc *services.MetricService) {
	r := app.Group("/metric")

	r.Get("/", svc.GetMetrics)
	r.Post("/", svc.AddMetric)
}

package metrics

import (
	"github.com/gofiber/fiber/v2"
)

// MetricRouter contains all the metrics routes
func MetricRouter(app fiber.Router, svc *MetricService) {
	r := app.Group("/metric")

	r.Get("/", svc.GetMetrics)
	r.Post("/", svc.AddMetric)
}

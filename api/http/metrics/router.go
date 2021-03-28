package metrics

import (
	"github.com/gofiber/fiber/v2"
)

// Router contains all the metrics routes
func Router(app fiber.Router, svc *Service) {
	r := app.Group("/metric")

	r.Get("/", svc.GetMetrics)
	r.Post("/", svc.AddMetric)
}

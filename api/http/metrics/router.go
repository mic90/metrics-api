package metrics

import (
	"github.com/gofiber/fiber/v2"
)

// Router contains all the metrics routes
func Router(app fiber.Router, svc *Service) {
	r := app.Group("/metrics")

	// metrics related endpoints
	r.Get("/", svc.GetMetrics)
	r.Post("/", svc.AddMetric)
	r.Delete("/:type/:name", svc.RemoveMetric)

	// metric data related endpoints
	r.Get("/:type/:name/data", svc.GetData)
	r.Post("/:type/:name/data", svc.AddData)
	r.Get("/:type/:name/data/:reducer", svc.GetDataReduced)
}

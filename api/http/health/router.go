package health

import (
	"github.com/gofiber/fiber/v2"
)

// Router contains all health routes
func Router(app fiber.Router, svc Service) {
	r := app.Group("/health")

	r.Get("/", svc.HealthCheck)
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mic90/metrics-api/api/http/services"
)

// HealthRouter contains all health routes
func HealthRouter(app fiber.Router, svc services.HealthService) {
	r := app.Group("/health")

	r.Get("/", svc.HealthCheck)
}

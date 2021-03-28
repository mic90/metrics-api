package health

import (
	"github.com/gofiber/fiber/v2"
)

// HealthRouter contains all health routes
func HealthRouter(app fiber.Router, svc HealthService) {
	r := app.Group("/health")

	r.Get("/", svc.HealthCheck)
}

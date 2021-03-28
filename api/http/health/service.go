package health

import (
	"github.com/gofiber/fiber/v2"
)

// HealthService is a dummy health service to indicate health status
type HealthService struct {
}

// HealthCheck returns service health
// As service doesn't relay on any external systems, it always returns OK
func (h HealthService) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(HealthCheck{
		Status: HealthOk,
	})
}

package health

import (
	"github.com/gofiber/fiber/v2"
)

// Service is a dummy health service to indicate health status
type Service struct {
}

// HealthCheck returns service health
// As service doesn't relay on any external systems, it always returns OK
func (h Service) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(HealthCheck{
		Status: HealthOk,
	})
}

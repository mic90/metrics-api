package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mic90/metrics-api/api/http/dto"
)

// HealthService is a dummy health service to indicate health status
type HealthService struct {
}

// HealthCheck returns service health
// As service doesn't relay on any external systems, it always returns OK
func (h HealthService) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(dto.HealthCheck{
		Status: dto.HealthOk,
	})
}

package health

import (
	"github.com/gofiber/fiber/v2"
)

// Service is a dummy health service to indicate health status
type Service struct {
}

// HealthCheck returns service health
// As service doesn't relay on any external systems, it always returns OK
// AddMetric adds new metric based on provided descriptor
// @Summary HealthCheck
// @Description returns health status of the serivce
// @Tags health
// @Produce  json
// @Success 200 {object} StatusResponse
// @Failure 503 {object} StatusResponse
// @Router /health [get]
func (h Service) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(StatusResponse{
		Status: OK,
	})
}

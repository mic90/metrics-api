package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/persistance"
)

// Service allows metrics manipulation
type Service struct {
	driver persistance.Storage
}

// NewService returns Service with specified storage driver
func NewService(driver persistance.Storage) *Service {
	return &Service{
		driver,
	}
}

// GetMetrics returns all metrics
func (m *Service) GetMetrics(ctx *fiber.Ctx) error {
	met := m.driver.GetMetrics()
	ret := make([]MetricDescriptor, 0, len(met))

	for _, value := range met {
		ret = append(ret, MetricDescriptor{
			Name: value.Name,
			Type: value.Type,
		})
	}

	return ctx.JSON(ret)
}

// GetMetrics returns all metrics
func (m *Service) AddMetric(ctx *fiber.Ctx) error {
	var descDTO MetricDescriptor
	if err := ctx.BodyParser(&descDTO); err != nil {
		return fiber.ErrBadRequest
	}

	descriptor := metrics.Descriptor{
		Name: descDTO.Name,
		Type: descDTO.Type,
	}
	if err := m.driver.AddMetric(descriptor); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

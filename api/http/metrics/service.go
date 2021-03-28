package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/persistance"
)

// MetricService allows metrics manipulation
type MetricService struct {
	driver persistance.Storage
}

// NewMetricService returns MetricService with specified storage driver
func NewMetricService(driver persistance.Storage) *MetricService {
	return &MetricService{
		driver,
	}
}

// GetMetrics returns all metrics
func (m *MetricService) GetMetrics(ctx *fiber.Ctx) error {
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
func (m *MetricService) AddMetric(ctx *fiber.Ctx) error {
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

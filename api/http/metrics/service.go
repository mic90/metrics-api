package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/persistance"
	"time"
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
// @Summary GetMetrics
// @Description returns all metrics descriptions
// @Tags metrics
// @Accept  json
// @Produce  json
// @Success 200 {array}  MetricDescriptor
// @Router /metrics [get]
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

// AddMetric adds new metric based on provided descriptor
// @Summary AddMetric
// @Description adds new metric based on provided descriptor
// @Tags metrics
// @Accept  json
// @Produce  json
// @Param message body MetricDescriptor true "Metric descriptor"
// @Success 200
// @Failure 400 {string} string "error"
// @Router /metrics [post]
func (m *Service) AddMetric(ctx *fiber.Ctx) error {
	var descDTO MetricDescriptor
	if err := ctx.BodyParser(&descDTO); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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

// RemoveMetric removes metric with all its data
// @Summary RemoveMetric
// @Description removes metric with all its data
// @Tags metrics
// @Param type path string true "Metric type"
// @Param name path string true "Metric name"
// @Success 200
// @Failure 400 {string} string "error"
// @Router /metrics/:type/:name [delete]
func (m *Service) RemoveMetric(ctx *fiber.Ctx) error {
	metricType := ctx.Params("type")
	metricName := ctx.Params("name")

	if metricType == "" || metricName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "metric description params are required")
	}

	desc := metrics.Descriptor{
		Name: metricName,
		Type: metricType,
	}

	if err := m.driver.RemoveMetric(desc); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

// AddData adds new data point to the metric
// @Summary AddData
// @Description adds new data point to the metric
// @Tags metrics
// @Accept  json
// @Produce  json
// @Param type path string true "Metric type"
// @Param name path string true "Metric name"
// @Param message body Value true "Metric descriptor with value"
// @Success 200
// @Failure 400 {string} string "error"
// @Router /metrics/:type/:name/data [post]
func (m *Service) AddData(ctx *fiber.Ctx) error {
	metricType := ctx.Params("type")
	metricName := ctx.Params("name")

	if metricType == "" || metricName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "metric description params are required")
	}

	var valueDTO Value
	if err := ctx.BodyParser(&valueDTO); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	descriptor := metrics.Descriptor{
		Name: metricName,
		Type: metricType,
	}
	if err := m.driver.AddData(descriptor, valueDTO.Value); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

// GetData returns data points for metric in given time range
// @Summary GetData
// @Description returns data points for metric in given time range
// @Tags metrics
// @Accept  json
// @Produce  json
// @Param type path string true "Metric type"
// @Param name path string true "Metric name"
// @Param from query string true "Begin timestamp in RFC3339 format"
// @Param to query string true "End timestamp in RFC3339 format"
// @Success 200 {array} MetricDataPoint
// @Failure 400 {string} string "Bad parameters provided by user"
// @Failure 500 {string} string "Data retrieval field"
// @Router /metrics/:type/:name/data [get]
func (m *Service) GetData(ctx *fiber.Ctx) error {
	metricType := ctx.Params("type")
	metricName := ctx.Params("name")

	if metricType == "" || metricName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "metric description params are required")
	}

	from := ctx.Query("from")
	to := ctx.Query("to")

	if from == "" || to == "" {
		return fiber.NewError(fiber.StatusBadRequest, "full date range is required")
	}

	var (
		fromTime   time.Time
		toTime     time.Time
		parseError error
	)

	if fromTime, parseError = time.Parse(time.RFC3339, from); parseError != nil {
		return fiber.NewError(fiber.StatusBadRequest, parseError.Error())
	}
	if toTime, parseError = time.Parse(time.RFC3339, to); parseError != nil {
		return fiber.NewError(fiber.StatusBadRequest, parseError.Error())
	}

	descriptor := metrics.Descriptor{
		Name: metricName,
		Type: metricType,
	}
	if data, err := m.driver.GetData(descriptor, fromTime, toTime); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else {
		return ctx.JSON(data)
	}
}

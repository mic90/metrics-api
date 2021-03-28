package persistance

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

type Storage interface {
	GetMetrics() []metrics.Descriptor
	AddMetric(metrics.Descriptor) error
	RemoveMetric(metrics.Descriptor) error

	AddData(metrics.Descriptor, float64) error
	GetData(metrics.Descriptor, time.Time, time.Time) ([]data.Point, error)
}

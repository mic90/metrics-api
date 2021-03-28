package persistance

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

type Storage interface {
	GetMetrics() []metrics.Metric
	AddMetric(metrics.Metric) error
	RemoveMetric(string) error

	AddData(string, float64) error
	GetData(string, time.Time, time.Time) ([]data.Point, error)
}

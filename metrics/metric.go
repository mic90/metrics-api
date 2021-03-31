package metrics

import (
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

// Timed is used for object with time range availability
type Timed interface {
	MinT() time.Time
	MaxT() time.Time
}

// Metric describes generic metric with time range and data points
type Metric interface {
	Timed

	AddData(data.Point) error
	Data() []data.Point
	Last() data.Point
	Size() int
}

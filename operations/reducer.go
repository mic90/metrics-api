package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"time"
)

// Reducer is an operation that reduces metric data points to single value
type Reducer interface {
	Process(metrics.Descriptor, time.Time, time.Time) (float64, error)
}

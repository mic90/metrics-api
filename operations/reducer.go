package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"time"
)

type Reducer interface {
	Process(metrics.Descriptor, time.Time, time.Time) (float64, error)
}

package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"time"
)

// Avg returns average value from metric data points
type Avg struct {
	storage persistance.Storage
}

// NewAvg returns new avg reducer
func NewAvg(storage persistance.Storage) Reducer {
	return &Avg{
		storage,
	}
}

// Process calculates average value
func (m Avg) Process(desc metrics.Descriptor, from, to time.Time) (value float64, err error) {
	var values []data.Point
	if values, err = m.storage.GetData(desc, from, to); err != nil {
		return value, err
	}

	for _, v := range values {
		value += v.Value
	}

	value /= float64(len(values))

	return value, nil
}

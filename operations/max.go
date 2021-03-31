package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"time"
)

// Max returns maximum value from metric data points
type Max struct {
	storage persistance.Storage
}

// NewMax returns new max reducer
func NewMax(storage persistance.Storage) Reducer {
	return &Max{
		storage,
	}
}

// Process calculates maximum value
func (m Max) Process(desc metrics.Descriptor, from, to time.Time) (value float64, err error) {
	var values []data.Point
	if values, err = m.storage.GetData(desc, from, to); err != nil {
		return value, err
	}

	for _, v := range values {
		if v.Value > value {
			value = v.Value
		}
	}

	return value, nil
}

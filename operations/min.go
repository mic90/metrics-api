package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"time"
)

// Min returns minimum value from metric data points
type Min struct {
	storage persistance.Storage
}

// NewMin returns new min reducer
func NewMin(storage persistance.Storage) Reducer {
	return &Min{
		storage,
	}
}

// Process calculates minimum value
func (m Min) Process(desc metrics.Descriptor, from, to time.Time) (value float64, err error) {
	var values []data.Point
	if values, err = m.storage.GetData(desc, from, to); err != nil {
		return value, err
	}

	for _, v := range values {
		if v.Value < value {
			value = v.Value
		}
	}

	return value, nil
}

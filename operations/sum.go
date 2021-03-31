package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"time"
)

// Sum returns sum of all data points from metric
type Sum struct {
	storage persistance.Storage
}

// NewSum returns new sum reducer
func NewSum(storage persistance.Storage) Reducer {
	return &Sum{
		storage,
	}
}

// Process calculates sum of all data points
func (s Sum) Process(desc metrics.Descriptor, from, to time.Time) (value float64, err error) {
	var values []data.Point
	if values, err = s.storage.GetData(desc, from, to); err != nil {
		return value, err
	}

	for _, v := range values {
		value += v.Value
	}

	return value, nil
}

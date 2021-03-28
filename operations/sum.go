package operations

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"time"
)

type Sum struct {
	storage persistance.Storage
}

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
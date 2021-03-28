package operations

import (
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"time"
)

type Max struct {
	storage persistance.Storage
}

func (m Max) Process(metric string, from, to time.Time) (value float64, err error) {
	var values []data.Point
	if values, err = m.storage.GetData(metric, from, to); err != nil {
		return value, err
	}

	for _, v := range values {
		if v.Value > value {
			value = v.Value
		}
	}

	return value, nil
}

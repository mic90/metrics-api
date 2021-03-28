package metrics

import (
	"errors"
	"github.com/mic90/metrics-api/metrics/data"
)

var ErrInvalidValue = errors.New("counter can be only increased or reset")

type Counter struct {
	rawMetric
}

func NewCounter() Metric {
	return &Counter{
		*newRaw(),
	}
}

func (c *Counter) AddData(dataPoint data.Point) error {
	if dataPoint.Value < c.lastValue.Value && dataPoint.Value != 0 {
		return ErrInvalidValue
	}

	c.values.Add(dataPoint)
	c.lastValue = dataPoint
	c.to = dataPoint.Time

	// if it's the first data point in metric, update metric start time
	if len(c.values.Data()) == 1 {
		c.from = dataPoint.Time
	}

	return nil
}

package metrics

import (
	"errors"
	"github.com/mic90/metrics-api/metrics/data"
)

type Counter struct {
	rawMetric
}

func NewCounter(name string) Metric {
	return &Counter{
		*newRaw(name),
	}
}

func (c *Counter) AddData(dataPoint data.Point) error {
	if dataPoint.Value < c.lastValue.Value {
		return errors.New("counter can't be decreased")
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

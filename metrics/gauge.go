package metrics

import (
	"github.com/mic90/metrics-api/metrics/data"
)

type Gauge struct {
	rawMetric
}

func NewGauge() Metric {
	return &Gauge{
		*newRaw(),
	}
}

func (c *Gauge) AddData(dataPoint data.Point) error {
	c.values.Add(dataPoint)
	c.lastValue = dataPoint
	c.to = dataPoint.Time

	// if it's the first data point in metric, update metric start time
	if len(c.values.Data()) == 1 {
		c.from = dataPoint.Time
	}

	return nil
}

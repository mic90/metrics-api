package metrics

import (
	"github.com/mic90/metrics-api/metrics/data"
)

type Gauge struct {
	rawMetric
}

func NewGauge(name string) Metric {
	return &Gauge{
		*newRaw(name),
	}
}

func (c *Gauge) AddData(dataPoint data.Point) error {
	c.values.Add(dataPoint)
	c.lastValue = dataPoint
	c.to = dataPoint.Time

	return nil
}

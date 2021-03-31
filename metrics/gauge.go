package metrics

import (
	"github.com/mic90/metrics-api/metrics/data"
)

// Gauge A gauge is a metric that represents a single numerical value that can arbitrarily go up and down.
type Gauge struct {
	rawMetric
}

// NewGauge creates a new gauge with default data buffer
func NewGauge() Metric {
	return &Gauge{
		*newRaw(),
	}
}

// AddData adds new data point to the metric
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

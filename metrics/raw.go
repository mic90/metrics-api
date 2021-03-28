package metrics

import (
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

type rawMetric struct {
	name      string
	values    data.Buffer
	lastValue data.Point
	from      time.Time
	to        time.Time
}

func newRaw(name string) *rawMetric {
	buf := data.NewBuffer()
	return &rawMetric{
		name:   name,
		values: *buf,
		lastValue: data.Point{
			Value: 0,
			Time:  time.Now(),
		},
		from: time.Now(),
		to:   time.Now(),
	}
}

func (c rawMetric) Name() string {
	return c.name
}

func (c rawMetric) Last() data.Point {
	return c.lastValue
}

func (c rawMetric) Data() []data.Point {
	return c.values.Data()
}

func (c rawMetric) Size() int {
	return len(c.values.Data())
}

func (c rawMetric) MinT() time.Time {
	return c.from
}

func (c rawMetric) MaxT() time.Time {
	return c.to
}

package metrics

import (
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

type Timed interface {
	MinT() time.Time
	MaxT() time.Time
}

type Metric interface {
	Timed

	Name() string

	AddData(data.Point) error
	Data() []data.Point
	Last() data.Point
	Size() int
}

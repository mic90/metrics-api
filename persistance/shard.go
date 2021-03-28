package persistance

import (
	"errors"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

type Shard struct {
	metric   metrics.Metric
	duration time.Duration
	end      time.Time
}

var ErrShardMaxReached = errors.New("reached maximum timestamp of the shard")

func NewShard(metric metrics.Metric, duration time.Duration) *Shard {
	return &Shard{
		metric,
		duration,
		metric.MinT().Add(duration),
	}
}

func (s *Shard) AddData(dataPoint data.Point) error {
	if dataPoint.Time.After(s.end) {
		return ErrShardMaxReached
	}

	return s.metric.AddData(dataPoint)
}

func (s *Shard) Data() []data.Point {
	return s.metric.Data()
}

func (s *Shard) DataFrom(t time.Time) []data.Point {
	for index, dp := range s.metric.Data() {
		if dp.Time.Before(t) {
			continue
		}

		return s.metric.Data()[index:]
	}

	return []data.Point{}
}

func (s *Shard) DataTo(t time.Time) []data.Point {
	for index, dp := range s.metric.Data() {
		if dp.Time.After(t) {
			return s.metric.Data()[:index]
		}
	}

	return s.metric.Data()
}

func (s Shard) Duration() time.Duration {
	return s.metric.MaxT().Sub(s.metric.MinT())
}

func (s Shard) EndT() time.Time {
	return s.end
}

func (s Shard) MinT() time.Time {
	return s.metric.MinT()
}

func (s Shard) MaxT() time.Time {
	return s.metric.MaxT()
}

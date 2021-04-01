package persistance

import (
	"errors"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

var ErrShardMaxReached = errors.New("reached maximum timestamp of the shard")

// Shard contains given metric data in specified time range
type Shard struct {
	metric   metrics.Metric
	duration time.Duration
	end      time.Time
}

// NewShard creates new shard based on provided metric and required max duration
func NewShard(metric metrics.Metric, duration time.Duration) *Shard {
	return &Shard{
		metric,
		duration,
		metric.MinT().Add(duration),
	}
}

// AddData adds new data point to the shard
func (s *Shard) AddData(dataPoint data.Point) error {
	if dataPoint.Time.After(s.end) {
		return ErrShardMaxReached
	}

	return s.metric.AddData(dataPoint)
}

// Data returns all data points
func (s *Shard) Data() []data.Point {
	return s.metric.Data()
}

// DataFrom returns data points starting from time t
func (s *Shard) DataFrom(t time.Time) []data.Point {
	for index, dp := range s.metric.Data() {
		if dp.Time.Before(t) {
			continue
		}

		return s.metric.Data()[index:]
	}

	return []data.Point{}
}

// DataTo returns data points up to time t
func (s *Shard) DataTo(t time.Time) []data.Point {
	for index, dp := range s.metric.Data() {
		if dp.Time.After(t) {
			return s.metric.Data()[:index]
		}
	}

	return s.metric.Data()
}

// DataRange returns data points in specified time range
func (s *Shard) DataRange(from, to time.Time) []data.Point {
	var (
		fromIndexFound bool
		fromIndex      int
		toIndex        int
	)

	for index, dp := range s.metric.Data() {
		if (dp.Time.Equal(from) || dp.Time.After(from)) && !fromIndexFound {
			fromIndex = index
			fromIndexFound = true
		}
		if dp.Time.After(to) {
			toIndex = index
			break
		}
	}

	// if there is no data that starts after start index, return empty array
	if !fromIndexFound {
		return []data.Point{}
	}

	return s.metric.Data()[fromIndex:toIndex]
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

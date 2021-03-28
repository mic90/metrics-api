package persistance_test

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShard_AddData(t *testing.T) {
	m, err := metrics.FromDescriptor(metricName, metricType)

	assert.NoError(t, err)

	s := persistance.NewShard(m, shardDuration)
	timestamp := time.Now()

	err = s.AddData(data.Point{
		Value: 10.0,
		Time:  timestamp,
	})

	assert.NoError(t, err)
	assert.Equal(t, timestamp, s.MinT())
	assert.Equal(t, timestamp, s.MaxT())
	assert.Equal(t, timestamp.Add(shardDuration), s.EndT())
}

func TestShard_Data_OnEmptyShard(t *testing.T) {
	m, err := metrics.FromDescriptor(metricName, metricType)

	assert.NoError(t, err)

	s := persistance.NewShard(m, shardDuration)
	d := s.Data()

	assert.Equal(t, 0, len(d))
}

func TestShard_Data(t *testing.T) {
	m, err := metrics.FromDescriptor(metricName, metricType)

	assert.NoError(t, err)

	s := persistance.NewShard(m, shardDuration)
	timestamp := time.Now()
	value := 10.0

	err = s.AddData(data.Point{
		Value: value,
		Time:  timestamp,
	})

	d := s.Data()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(d))
	assert.Equal(t, 10.0, d[0].Value)
}

func TestShard_DataFrom(t *testing.T) {
	m, err := metrics.FromDescriptor(metricName, metricType)

	assert.NoError(t, err)

	s := persistance.NewShard(m, shardDuration)
	timestamp := time.Now()
	value := 10.0

	err = s.AddData(data.Point{
		Value: value,
		Time:  timestamp,
	})

	d := s.DataFrom(timestamp.Add(shardDuration))

	assert.NoError(t, err)
	assert.Equal(t, 0, len(d))
}

func TestShard_DataTo(t *testing.T) {
	m, err := metrics.FromDescriptor(metricName, metricType)

	assert.NoError(t, err)

	s := persistance.NewShard(m, shardDuration)
	timestamp := time.Now()
	value := 10.0

	err = s.AddData(data.Point{
		Value: value,
		Time:  timestamp,
	})

	d := s.DataTo(timestamp.Add(-shardDuration))

	assert.NoError(t, err)
	assert.Equal(t, 0, len(d))
}

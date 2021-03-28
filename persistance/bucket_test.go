package persistance_test

import (
	"github.com/bradfitz/iter"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	shardDuration = 1 * time.Minute
	metricType    = "counter"
	metricName    = "metric"
)

func TestBucket_AddData(t *testing.T) {
	b, err := persistance.NewBucket(metricName, metricType, shardDuration)

	assert.NoError(t, err)

	err = b.AddData(data.FromValue(10))

	assert.NoError(t, err)
	assert.Equal(t, 1, b.Size())
}

func TestBucket_AddData_ShouldCreateShard(t *testing.T) {
	b, err := persistance.NewBucket(metricName, metricType, shardDuration)

	assert.NoError(t, err)

	err = b.AddData(data.FromValue(10))

	assert.NoError(t, err)

	err = b.AddData(data.Point{
		Value: 15,
		Time:  time.Now().Add(2 * shardDuration),
	})

	assert.NoError(t, err)
	assert.Equal(t, 2, b.Size())
}

func TestBucket_Data_OnEmptyBucket(t *testing.T) {
	b, err := persistance.NewBucket(metricName, metricType, shardDuration)

	assert.NoError(t, err)

	d, err := b.Data(time.Now(), time.Now())

	assert.NoError(t, err)
	assert.Equal(t, 0, len(d))
}

func TestBucket_Data_OnMultipleShards(t *testing.T) {
	b, err := persistance.NewBucket(metricName, metricType, shardDuration)

	assert.NoError(t, err)

	startTime := time.Now()
	endTime := time.Now()
	dataCount := 10

	for i := range iter.N(dataCount) {
		err := b.AddData(data.Point{
			Value: float64(i),
			Time:  endTime,
		})

		assert.NoError(t, err)

		// make sure each data point will end up in different shards
		endTime = endTime.Add(shardDuration * 2)
	}

	// read whole time range
	d, err := b.Data(startTime, endTime)

	assert.NoError(t, err)
	assert.Equal(t, dataCount, b.Size())
	assert.Equal(t, dataCount, len(d))
}

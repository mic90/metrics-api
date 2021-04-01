package persistance_test

import (
	"github.com/bradfitz/iter"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	shardDuration = 1 * time.Minute
)

var metricDesc = metrics.Descriptor{
	Name: "metric",
	Type: "counter",
}

func TestBucket_AddData(t *testing.T) {
	b, err := persistance.NewBucket(metricDesc, shardDuration)

	assert.NoError(t, err)

	err = b.AddData(data.FromValue(10))

	assert.NoError(t, err)
	assert.Equal(t, 1, b.Size())
}

func TestBucket_AddData_ShouldCreateShard(t *testing.T) {
	b, err := persistance.NewBucket(metricDesc, shardDuration)

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
	b, err := persistance.NewBucket(metricDesc, shardDuration)

	assert.NoError(t, err)

	d := b.Data(time.Now(), time.Now())

	assert.Equal(t, 0, len(d))
}

func TestBucket_Data_OnMultipleShards(t *testing.T) {
	b, err := persistance.NewBucket(metricDesc, shardDuration)

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
		endTime = endTime.Add(shardDuration)
	}

	// read whole time range
	d := b.Data(startTime, endTime)

	assert.True(t, b.Size() > 1)
	assert.Equal(t, dataCount, len(d))
	assert.Equal(t, float64(0), d[0].Value)
	assert.Equal(t, float64(dataCount-1), d[len(d)-1].Value)
}

func TestBucket_Data_EndAfterAvailableTimeRange(t *testing.T) {
	b, err := persistance.NewBucket(metricDesc, shardDuration)

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
		endTime = endTime.Add(shardDuration)
	}

	// from range should start somewhere in the middle of the data points
	from := startTime.Add(4 * time.Minute)
	// end time is long after last data point
	to := endTime.Add(1 * time.Hour)

	// read whole time range
	d := b.Data(from, to)

	assert.True(t, b.Size() > 1)
	assert.Equal(t, 6, len(d))
	assert.Equal(t, float64(4), d[0].Value)
	assert.Equal(t, float64(dataCount-1), d[len(d)-1].Value)
}

func TestBucket_Data_StartTimeAfterAvailableTimeRange(t *testing.T) {
	b, err := persistance.NewBucket(metricDesc, shardDuration)

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
		endTime = endTime.Add(shardDuration)
	}

	// time ranges are long after last data point
	from := startTime.Add(1 * time.Hour)
	to := from.Add(1 * time.Hour)

	d := b.Data(from, to)

	assert.True(t, b.Size() > 1)
	assert.Equal(t, 0, len(d))
}

func TestBucket_Data_OnOneShard(t *testing.T) {
	duration := 1 * time.Hour
	b, err := persistance.NewBucket(metricDesc, duration)

	assert.NoError(t, err)

	startTime := time.Now()
	endTime := time.Now()
	dataCount := 20

	for i := range iter.N(dataCount) {
		err := b.AddData(data.Point{
			Value: float64(i),
			Time:  endTime,
		})

		assert.NoError(t, err)

		// make sure each data point will end up in different shards
		endTime = endTime.Add(shardDuration)
	}

	// from range should start somewhere in the middle of the data points
	from := startTime.Add(4 * time.Minute)
	// end time should end somewhere before data range
	to := from.Add(5 * time.Minute)

	d := b.Data(from, to)

	assert.Equal(t, 1, b.Size())
	assert.Equal(t, 5, len(d))
	assert.Equal(t, 4.0, d[0].Value)
	assert.Equal(t, 8.0, d[len(d)-1].Value)
}

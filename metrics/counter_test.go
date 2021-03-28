package metrics_test

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter_AddValue(t *testing.T) {
	gauge := metrics.NewCounter("metric")
	value := data.FromValue(10)

	err := gauge.AddData(value)

	assert.NoError(t, err)
	assert.Equal(t, value, gauge.Last().Value)
}

func TestCounter_AddValue_ShouldFailOnDecrease(t *testing.T) {
	gauge := metrics.NewCounter("metric")
	value := data.FromValue(10)
	valueSec := data.FromValue(5)

	err := gauge.AddData(value)
	assert.NoError(t, err)

	err = gauge.AddData(valueSec)

	assert.Error(t, err)
	assert.Equal(t, value, gauge.Last().Value)
}

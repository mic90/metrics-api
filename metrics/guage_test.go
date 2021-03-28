package metrics_test

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGauge_AddValue(t *testing.T) {
	gauge := metrics.NewGauge("metric")
	value := data.FromValue(10)

	err := gauge.AddData(value)

	assert.NoError(t, err)
	assert.Equal(t, value, gauge.Last().Value)
}

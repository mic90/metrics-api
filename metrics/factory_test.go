package metrics_test

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromDescriptor_UnsupportedType(t *testing.T) {
	m, err := metrics.FromDescriptor(metrics.Descriptor{
		Name: "metric",
		Type: "wrongType",
	})

	assert.Error(t, err, metrics.ErrUnsupportedMetricType)
	assert.Nil(t, m)
}

func TestFromDescriptor_Counter(t *testing.T) {
	m, err := metrics.FromDescriptor(metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	})

	assert.NoError(t, err)
	assert.NotNil(t, m)
}

func TestFromDescriptor_Gauge(t *testing.T) {
	m, err := metrics.FromDescriptor(metrics.Descriptor{
		Name: "metric",
		Type: "gauge",
	})

	assert.NoError(t, err)
	assert.NotNil(t, m)
}

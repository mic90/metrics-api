package metrics_test

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDescriptor_Hash(t *testing.T) {
	desc := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}

	assert.Equal(t, "metriccounter", desc.Hash())
}

package data_test

import (
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

const bufferIncreaseStep = 1000

func TestBuffer_Add(t *testing.T) {
	buff := data.NewBuffer()

	buff.Add(data.FromValue(10))

	assert.Equal(t, bufferIncreaseStep, buff.Cap())
}

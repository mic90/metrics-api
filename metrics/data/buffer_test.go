package data_test

import (
	"github.com/bradfitz/iter"
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

func TestBuffer_Add_ShouldGrowAutomatically(t *testing.T) {
	buff := data.NewBuffer()

	for i := range iter.N(bufferIncreaseStep) {
		buff.Add(data.FromValue(float64(i)))
	}

	assert.Equal(t, bufferIncreaseStep*2, buff.Cap())
}

func TestBuffer_Grow(t *testing.T) {
	buff := data.NewBuffer()
	buff.Grow(bufferIncreaseStep * 2)

	buff.Add(data.FromValue(10))

	assert.Equal(t, bufferIncreaseStep*2, buff.Cap())
}

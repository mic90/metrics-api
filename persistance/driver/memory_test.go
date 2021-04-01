package driver_test

import (
	"github.com/bradfitz/iter"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/persistance/driver"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemory_GetMetrics(t *testing.T) {
	m := driver.NewMemory()
	metricDescriptor := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}

	err := m.AddMetric(metricDescriptor)
	ret := m.GetMetrics()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(ret))
	assert.Equal(t, metricDescriptor.Name, ret[0].Name)
	assert.Equal(t, metricDescriptor.Type, ret[0].Type)
}

func TestMemory_AddMetric_AlreadyExist(t *testing.T) {
	m := driver.NewMemory()
	metricDescriptor := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}

	err1 := m.AddMetric(metricDescriptor)
	err2 := m.AddMetric(metricDescriptor)

	assert.NoError(t, err1)
	assert.Error(t, err2, driver.ErrMetricAlreadyExists)
}

func TestMemory_RemoveMetric(t *testing.T) {
	m := driver.NewMemory()
	metricDescriptor := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}

	errAdd := m.AddMetric(metricDescriptor)
	errRemove := m.RemoveMetric(metricDescriptor)

	ret := m.GetMetrics()

	assert.NoError(t, errAdd)
	assert.NoError(t, errRemove)
	assert.Equal(t, 0, len(ret))
}

func TestMemory_AddData(t *testing.T) {
	m := driver.NewMemory()
	metricDescriptor := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}

	errAdd := m.AddMetric(metricDescriptor)
	errData := m.AddData(metricDescriptor, 10.0)

	assert.NoError(t, errAdd)
	assert.NoError(t, errData)
}

func TestMemory_AddData_MetricDoesNotExists(t *testing.T) {
	m := driver.NewMemory()
	metricDescriptor := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}
	nonExistingMetricDescriptor := metrics.Descriptor{
		Name: "metric2",
		Type: "counter",
	}

	errAdd := m.AddMetric(metricDescriptor)
	errData := m.AddData(nonExistingMetricDescriptor, 10.0)

	assert.NoError(t, errAdd)
	assert.Error(t, errData, driver.ErrMetricDoesNotExists)
}

func TestMemory_GetData(t *testing.T) {
	m := driver.NewMemory()
	metricDescriptor := metrics.Descriptor{
		Name: "metric",
		Type: "counter",
	}
	dataPointsCount := 100

	errAdd := m.AddMetric(metricDescriptor)

	assert.NoError(t, errAdd)

	for i := range iter.N(dataPointsCount) {
		if err := m.AddData(metricDescriptor, float64(i)); err != nil {
			t.Fatal(err)
		}
	}

	// make sure all data points are in the time range
	from := time.Now().Add(-1 * time.Minute)
	to := from.Add(1 * time.Minute)

	ret, err := m.GetData(metricDescriptor, from, to)
	assert.NoError(t, err)
	assert.Equal(t, dataPointsCount, len(ret))
	assert.Equal(t, float64(0), ret[0].Value)
	assert.Equal(t, float64(dataPointsCount-1), ret[len(ret)-1].Value)
}

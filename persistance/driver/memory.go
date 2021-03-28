package driver

import (
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/persistance"
	"sync"
	"time"
)

const shardDuration = 30 * time.Minute

type Memory struct {
	mutex   sync.RWMutex
	metrics map[string]*persistance.Bucket
}

func NewMemory() persistance.Storage {
	return &Memory{
		metrics: make(map[string]*persistance.Bucket),
	}
}

func (m *Memory) GetMetrics() []metrics.Descriptor {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	ret := make([]metrics.Descriptor, 0, len(m.metrics))

	for _, bucket := range m.metrics {
		ret = append(ret, bucket.Descriptor())
	}

	return ret
}

func (m *Memory) AddMetric(desc metrics.Descriptor) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.metrics[desc.Hash()]; ok {
		return ErrMetricAlreadyExists
	}

	if bucket, err := persistance.NewBucket(desc, shardDuration); err != nil {
		return err
	} else {
		m.metrics[desc.Hash()] = bucket
	}

	return nil
}

func (m *Memory) RemoveMetric(desc metrics.Descriptor) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.metrics, desc.Hash())

	return nil
}

func (m *Memory) AddData(desc metrics.Descriptor, value float64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if bucket, ok := m.metrics[desc.Hash()]; !ok {
		return ErrMetricDoesNotExists
	} else if err := bucket.AddData(data.FromValue(value)); err != nil {
		return err
	}

	return nil
}

func (m *Memory) GetData(desc metrics.Descriptor, from, to time.Time) ([]data.Point, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var ret []data.Point

	if bucket, ok := m.metrics[desc.Hash()]; !ok {
		return ret, ErrMetricDoesNotExists
	} else {
		ret = bucket.Data(from, to)
	}

	return ret, nil
}

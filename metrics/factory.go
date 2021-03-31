package metrics

import "errors"

var ErrUnsupportedMetricType = errors.New("unsupported metric type")

// FromDescriptor creates metric based on descriptor
func FromDescriptor(desc Descriptor) (Metric, error) {
	switch {
	case desc.Type == "counter":
		return NewCounter(), nil
	case desc.Type == "gauge":
		return NewGauge(), nil
	default:
		return nil, ErrUnsupportedMetricType
	}
}

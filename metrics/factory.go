package metrics

import "errors"

var ErrUnsupportedMetricType = errors.New("unsupported metric type")

func FromType(name, _type string) (Metric, error) {
	switch {
	case _type == "counter":
		return NewCounter(name), nil
	case _type == "gauge":
		return NewGauge(name), nil
	default:
		return nil, ErrUnsupportedMetricType
	}
}

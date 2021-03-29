package metrics

import "time"

type MetricDescriptor struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Value struct {
	Value float64 `json:"value"`
}

type MetricDataPoint struct {
	Value
	Time time.Time `json:"time"`
}

type MetricDataPoints []MetricDataPoint

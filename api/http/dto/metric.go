package dto

type MetricDescriptor struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type MetricDescriptorValued struct {
	MetricDescriptor
	Value float64 `json:"value"`
}

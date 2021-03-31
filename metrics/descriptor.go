package metrics

import "fmt"

// Descriptor contains basic metric data that uniquely identifies given metric
type Descriptor struct {
	Name string
	Type string
}

// Hash returns metric hash value, which should be unique in app space
func (d Descriptor) Hash() string {
	return fmt.Sprintf("%s%s", d.Name, d.Type)
}

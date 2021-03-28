package metrics

import "fmt"

type Descriptor struct {
	Name string
	Type string
}

func (d Descriptor) Hash() string {
	return fmt.Sprintf("%s%s", d.Name, d.Type)
}

package operations

import "time"

type Reducer interface {
	Process(string, time.Time, time.Time) (float64, error)
}

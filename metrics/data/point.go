package data

import "time"

type Point struct {
	Value float64
	Time  time.Time
}

func FromValue(value float64) Point {
	return Point{
		value,
		time.Now(),
	}
}

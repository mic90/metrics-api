package operations

import (
	"errors"
	"github.com/mic90/metrics-api/persistance"
)

var ErrUnsupportedReducerName = errors.New("unsupported reducer name")

// FromName creates reducer based on it's name
func FromName(name string, storage persistance.Storage) (Reducer, error) {
	switch name {
	case "sum":
		return NewSum(storage), nil
	case "avg":
		return NewAvg(storage), nil
	case "min":
		return NewMin(storage), nil
	case "max":
		return NewMax(storage), nil
	default:
		return nil, ErrUnsupportedReducerName
	}
}

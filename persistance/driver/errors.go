package driver

import "errors"

var ErrMetricAlreadyExists = errors.New("metric already exists")
var ErrMetricDoesNotExists = errors.New("metric does not exists")

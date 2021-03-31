package operations_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"github.com/mic90/metrics-api/operations"
	mock_persistance "github.com/mic90/metrics-api/persistance/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAvg_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := []data.Point{
		data.FromValue(0.0),
		data.FromValue(5.0),
		data.FromValue(10.0),
		data.FromValue(15.0),
		data.FromValue(20.0),
	}
	expectedAvg := 10.0

	storage := mock_persistance.NewMockStorage(ctrl)
	storage.EXPECT().GetData(gomock.Any(), gomock.Any(), gomock.Any()).Return(d, nil)

	op := operations.NewAvg(storage)
	ret, err := op.Process(metrics.Descriptor{Name: "metric", Type: "counter"}, time.Now(), time.Now())

	assert.NoError(t, err)
	assert.Equal(t, expectedAvg, ret)
}

func TestAvg_Process_StorageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_persistance.NewMockStorage(ctrl)
	storage.EXPECT().GetData(gomock.Any(), gomock.Any(), gomock.Any()).Return([]data.Point{}, errors.New("storage error"))

	op := operations.NewAvg(storage)
	ret, err := op.Process(metrics.Descriptor{Name: "metric", Type: "counter"}, time.Now(), time.Now())

	assert.Error(t, err)
	assert.Equal(t, 0.0, ret)
}

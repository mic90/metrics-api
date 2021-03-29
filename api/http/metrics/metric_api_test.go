package metrics_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	api "github.com/mic90/metrics-api/api/http"
	"github.com/mic90/metrics-api/api/http/health"
	"github.com/mic90/metrics-api/api/http/metrics"
	mocks "github.com/mic90/metrics-api/persistance/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMetrics_AddMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().AddMetric(gomock.Any()).Return(nil)

	healthSvc := &health.Service{}
	metricsSvc := metrics.NewService(storage)

	app := api.Setup(healthSvc, metricsSvc)

	desc := metrics.MetricDescriptor{
		Name: "metric",
		Type: "counter",
	}
	reqBody, err := json.Marshal(&desc)
	assert.NoError(t, err)

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/metrics",
		bytes.NewReader(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestMetrics_AddMetrics_UnsupportedMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	storage.EXPECT().AddMetric(gomock.Any()).Return(errors.New("unsupported metric type"))

	healthSvc := &health.Service{}
	metricsSvc := metrics.NewService(storage)

	app := api.Setup(healthSvc, metricsSvc)

	desc := metrics.MetricDescriptor{
		Name: "metric",
		Type: "unsupportedMetricType",
	}
	reqBody, err := json.Marshal(&desc)
	assert.NoError(t, err)

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/metrics",
		bytes.NewReader(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, 400, res.StatusCode)
}
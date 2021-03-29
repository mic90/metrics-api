package health_test

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	api "github.com/mic90/metrics-api/api/http"
	"github.com/mic90/metrics-api/api/http/health"
	"github.com/mic90/metrics-api/api/http/metrics"
	mocks "github.com/mic90/metrics-api/persistance/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHealth_GetHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorage(ctrl)
	healthSvc := &health.Service{}
	metricsSvc := metrics.NewService(storage)

	app := api.Setup(healthSvc, metricsSvc)

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/health",
		nil,
	)

	res, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var status health.StatusResponse
	err = json.Unmarshal(body, &status)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, health.OK, status.Status)
}
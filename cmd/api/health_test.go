package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	app := newTestApplication(t);
	router := app.mount()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := jsonResponseEnvelopeStrict[healthCheckResponse]{
		Data : healthCheckResponse{
			Status:  "ok",
			Version: "0.0.0",
			Env:     "test-env",
		},
	}

	var actualResponse jsonResponseEnvelopeStrict[healthCheckResponse]
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse, actualResponse)
}


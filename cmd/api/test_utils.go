package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()
	gin.SetMode(gin.TestMode)

	mockTodos := new(store.MockTodos)

	app := &application{
		config: config{
			version: "0.0.0",
			env:     "test-env",
		},
		store: store.Storage{
			Todos: mockTodos,
		},
	}

	return app
}
func executeRequest(t *testing.T, router http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	t.Helper()

	w := httptest.NewRecorder()
	var reqBody io.Reader

	switch v := body.(type) {
	case nil:
			reqBody = nil
	case string:
			reqBody = strings.NewReader(v)
	case []byte:
			reqBody = bytes.NewReader(v)
	default:
			jsonBytes, err := json.Marshal(v)
			if err != nil {
					t.Fatalf("failed to marshal request body: %v", err)
			}
			reqBody = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest(method, path, reqBody)
	if err != nil {
			t.Fatalf("failed to create request: %v", err)
	}

	router.ServeHTTP(w, req)
	return w
}

func checkJSONResponse[T any](t *testing.T, w *httptest.ResponseRecorder, expected T) {
	t.Helper() 
	var actualResponse T
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)

	assert.Equal(t, expected, actualResponse)
}


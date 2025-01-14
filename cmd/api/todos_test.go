package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestGetTodoByID(t *testing.T) {

	app := newTestApplication(t);
	router := app.mount()

	mockTodos := app.store.Todos.(*store.MockTodos)

	t.Run("successfully", func(t *testing.T) {

		expectedTodo := &store.Todo{
			ID:        1,
			Item:      "Test Todo",
			Completed: false,
			CreatedAt: "",  
			UpdatedAt: "",
			UserID:    0,
		}
	
		mockTodos.On("GetByID", context.Background(), int64(1)).Return(expectedTodo, nil)

		w := executeRequest(t, router, "GET", "/v1/todos/1", nil)

		assert.Equal(t, http.StatusOK, w.Code)

		// TODO: Left it only for an alternative way to check response
		// checkJSONResponse(t, w, expectedResponse) is much better
		// see "TestCreateTodo successfuly" test case below

		expectedResponse := `{
			"data": {
				"id": 1,
				"item": "Test Todo",
				"completed": false,
				"created_at": "",
				"updated_at": "",
				"user_id": 0
			}
		}`
		assert.JSONEq(t, expectedResponse, w.Body.String())

		mockTodos.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {

		w := executeRequest(t, router, "GET", "/v1/todos/qwe", nil)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "query param validation failed")
	})

	t.Run("not found", func(t *testing.T) {

		mockTodos.On("GetByID", context.Background(), int64(99)).Return(&store.Todo{}, store.ErrNotFound)

		w := executeRequest(t, router, "GET", "/v1/todos/99", nil)

		assert.Equal(t, http.StatusNotFound, w.Code)

		expectedResponse := `{
			"error": "not found"
		}`
		assert.JSONEq(t, expectedResponse, w.Body.String())

		mockTodos.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {

	app := newTestApplication(t);
	router := app.mount()

	mockTodos := app.store.Todos.(*store.MockTodos)

	t.Run("request body is empty", func(t *testing.T) {
		w := executeRequest(t, router, "POST", "/v1/todos", "")

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "request body is empty")
	})

	t.Run("bad request, Item is required", func(t *testing.T) {
		w := executeRequest(t, router, "POST", "/v1/todos", "{}")

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "json param validation failed", "Item", "required")
	})

	t.Run("bad request, Item is too short", func(t *testing.T) {

		payload := todoCreatePayload{
			Item: "w",
		}
		w := executeRequest(t, router, "POST", "/v1/todos", payload)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "json param validation failed", "Item", "min")
	})

	t.Run("bad request, Item is too long", func(t *testing.T) {
		payload := todoCreatePayload{
			Item: "This is the item with more than 100 characters. It is used to demonstrate how to create an item that exceeds the 100 character limit and still maintains clarity and readability.",
		}
		w := executeRequest(t, router, "POST", "/v1/todos", payload)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "json param validation failed", "Item", "max")
	})

	t.Run("bad request, Completed is not a boolean", func(t *testing.T) {
		w := executeRequest(t, router, "POST", "/v1/todos", `{"item":"Item","completed":"df"}`)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "json param validation failed", "completed",)
	})

	t.Run("successfuly", func(t *testing.T) {
    expectedResponse := jsonResponseEnvelopeStrict[store.Todo]{
        Data: store.Todo{
            ID:        0,
            Item:      "Test Todo",
            Completed: false,
            CreatedAt: "",  
            UpdatedAt: "",
            UserID:    1, 
        },
    }

    payload := todoCreatePayload{
        Item:      "Test Todo",  
        Completed: false,      
    }

    mockTodos.On("Create", context.Background(), &store.Todo{
        Item:      payload.Item,
        Completed: payload.Completed,
        UserID:    1,
    }).Return(nil)  

    w := executeRequest(t, router, "POST", "/v1/todos", payload)

    assert.Equal(t, http.StatusOK, w.Code)

    checkJSONResponse(t, w, expectedResponse)

    mockTodos.AssertExpectations(t)
})

}


package main

import (
	"strconv"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/cmd/internal/store"
	"github.com/gin-gonic/gin"
)

const todoCtx = "todo"

type todoCreatePayload struct {
	Item      string `json:"item" binding:"required,max=100,min=2"`
	Completed bool   `json:"completed"`
}

func (app *application) getTodosHandler(c *gin.Context) {
	todos, err := app.store.Todos.GetTodos(c.Request.Context())
	if err != nil {
		app.internalServerError(c)
		return
	}

	app.jsonOkResponse(c, todos)
}

func (app *application) createTodoHandler(c *gin.Context) {

	var payload todoCreatePayload

	if err := app.shouldBindJSON(c, &payload); err != nil {
		app.badRequestResponse(c, err)
		return
	}

	todo := &store.Todo{
		Item: payload.Item,
		Completed: payload.Completed,
	}

	ctx := c.Request.Context()

	if err := app.store.Todos.Create(ctx, todo); err != nil {
		app.internalServerError(c)
		return
	}

	app.jsonOkResponse(c, todo)
}

type todoUpdatePayload struct {
	Item      *string `json:"item" binding:"omitempty,max=100,min=2"`
	Completed *bool   `json:"completed" binding:"omitempty"`
}

func (app *application) updateTodoHandler(c *gin.Context) {

	todo := getTodoFromCtx(c)

	var payload todoUpdatePayload

	if err := app.shouldBindJSON(c, &payload); err != nil {
		app.badRequestResponse(c, err)
		return
	}

	if payload.Item != nil{
		todo.Item = *payload.Item
	}
	
	if payload.Completed != nil {
		todo.Completed = *payload.Completed
	}

	ctx := c.Request.Context()

	err := app.store.Todos.Update(ctx, todo);
	if err != nil {
		app.notFoundStoreResponse(c, err)
		return
	}

	app.jsonOkResponse(c, todo)
}

func (app *application) deleteTodoHandler(c *gin.Context) {
	todo := getTodoFromCtx(c)

	err := app.store.Todos.Delete(c.Request.Context(), todo.ID)
	if err != nil {
		app.notFoundStoreResponse(c, err)
		return
	}

	app.jsonOkResponse(c, todo)
}

func (app *application) getTodoHandler(c *gin.Context) {
	app.jsonOkResponse(c, getTodoFromCtx(c))
}

func getTodoFromCtx(c *gin.Context) *store.Todo {
	return c.MustGet(todoCtx).(*store.Todo)
}

type getTodoPayload struct {
	ID string `uri:"id" binding:"required,numericString"`
}

func (app *application) todoContextMiddleware() gin.HandlerFunc{
	return func (c *gin.Context) {
		var payload getTodoPayload

		if err := app.shouldBindUri(c, &payload); err != nil {
			app.badRequestResponse(c, err)
			return 
		}

		todoID, _ := strconv.ParseInt(payload.ID, 10, 64)

		ctx := c.Request.Context()

		todo, err := app.store.Todos.GetByID(ctx, todoID)
		if err != nil {
			app.notFoundStoreResponse(c, err)
			return
		}
		c.Set(todoCtx, todo)
		c.Next()
	}
}
package main

import (
	"strconv"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/store"
	"github.com/gin-gonic/gin"
)

const todoCtx = "todo"

type todoCreatePayload struct {
	Item      string `json:"item" binding:"required,max=100,min=2"`
	Completed bool   `json:"completed"`
}

// getTodosHandler godoc
//
//	@Summary		Get Todos
//	@Description	Returns a list of todos based on query parameters such as limit, offset, order, and search.
//	@Accept			json
//	@Tags			Todos
//	@Produce		json
//	@Param			limit	query		int		false	"Number of todos to return (min: 1, max: 100)"	default(10)
//	@Param			offset	query		int		false	"Number of todos to skip (min: 0)"				default(0)
//	@Param			order	query		string	false	"Order of the todos (asc or desc)"				default(desc)
//	@Param			search	query		string	false	"Search term to filter todos by title or description (max length: 100)"
//	@Success		200		{object}	jsonResponseEnvelope{data=[]store.Todo}
//	@Failure		400		{object}	jsonErrorResponseEnvelope
//	@Failure		500		{object}	jsonErrorResponseEnvelope
//	@Router			/todos [get]
func (app *application) getTodosHandler(c *gin.Context) {

	payload := store.TodosQueryParams{
		Limit: 10, 
		Order: "desc",
	}

	if err := app.shouldBindQuery(c, &payload); err != nil {
		app.badRequestResponse(c, err)
		return
	}

	todos, err := app.store.Todos.GetTodos(c.Request.Context(), payload)
	if err != nil {
		app.internalServerError(c)
		return
	}

	app.jsonOkResponse(c, todos)
}

// createTodoHandler godoc
//
//	@Summary		Create Todo
//	@Description	Creates a new todo item.
//	@Accept			json
//	@Tags			Todos
//	@Produce		json
//	@Param			payload	body		todoCreatePayload	true	"Todo item to be created"
//	@Success		200		{object}	jsonResponseEnvelope{data=store.Todo}
//	@Failure		400		{object}	jsonErrorResponseEnvelope
//	@Failure		500		{object}	jsonErrorResponseEnvelope
//	@Router			/todos [post]
func (app *application) createTodoHandler(c *gin.Context) {

	var payload todoCreatePayload

	if err := app.shouldBindJSON(c, &payload); err != nil {
		app.badRequestResponse(c, err)
		return
	}

	todo := &store.Todo{
		Item: payload.Item,
		Completed: payload.Completed,
		// TODO: get user id from token
		UserID: 1,
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

// updateTodoHandler godoc
//
//	@Summary		Update Todo
//	@Description	Updates an existing todo item by its ID.
//	@Tags			Todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int64				true	"ID of the todo to update"
//	@Param			body	body		todoUpdatePayload	true	"Todo update payload"
//	@Success		200		{object}	jsonResponseEnvelope{data=store.Todo}
//	@Failure		400		{object}	jsonErrorResponseEnvelope
//	@Failure		404		{object}	jsonErrorResponseEnvelope
//	@Failure		500		{object}	jsonErrorResponseEnvelope
//	@Router			/todos/{id} [put]
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

// deleteTodoHandler godoc
//
//	@Summary		Delete Todo
//	@Description	Deletes a todo item by its ID.
//	@Tags			Todos
//	@Produce		json
//	@Param			id	path		int64	true	"ID of the todo to delete"
//	@Success		200	{object}	jsonResponseEnvelope{data=store.Todo}
//	@Failure		404	{object}	jsonErrorResponseEnvelope
//	@Failure		500	{object}	jsonErrorResponseEnvelope
//	@Router			/todos/{id} [delete]
func (app *application) deleteTodoHandler(c *gin.Context) {
	todo := getTodoFromCtx(c)

	err := app.store.Todos.Delete(c.Request.Context(), todo.ID)
	if err != nil {
		app.notFoundStoreResponse(c, err)
		return
	}

	app.jsonOkResponse(c, todo)
}

// getTodoHandler godoc
//
//	@Summary		Get Todo
//	@Description	Returns a todo item by its ID.
//	@Tags			Todos
//	@Produce		json
//	@Param			id	path		int64	true	"ID of the todo to retrieve"
//	@Success		200	{object}	jsonResponseEnvelope{data=store.Todo}
//	@Failure		404	{object}	jsonErrorResponseEnvelope
//	@Failure		500	{object}	jsonErrorResponseEnvelope
//	@Router			/todos/{id} [get]
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
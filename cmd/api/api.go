package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) mount() http.Handler {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(app.errorHandler())

	v1 := r.Group("/v1")
	{
		v1.GET("/health", app.healthCheckHandler)

		todos := v1.Group("/todos") 
		{
			todos.GET("", app.getTodosHandler)
			todos.POST("", app.createTodoHandler)
	
			todo := todos.Group("/:id")
			{
				todo.Use(app.todoContextMiddleware())
		
				todo.GET("", app.getTodoHandler)
				todo.PATCH("", app.updateTodoHandler)
				todo.DELETE("", app.deleteTodoHandler)
			}
		}
	}

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
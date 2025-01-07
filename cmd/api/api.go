package main

import (
	"log"
	"net/http"
	"time"

	docs "github.com/aleksandr-slobodian/go-todolist-crud-postgres/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func (app *application) mount() http.Handler {
	r := gin.New()
	docs.SwaggerInfo.Version = app.config.version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(app.errorHandler())

	v1 := r.Group("/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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
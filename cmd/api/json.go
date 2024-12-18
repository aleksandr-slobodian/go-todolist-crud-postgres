package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) jsonResponse(c *gin.Context, status int, data any) {
	type envelope struct {
		Data any `json:"data"`
	}

	c.JSON(status, envelope{Data: data})
}

func (app *application) jsonOkResponse(c *gin.Context, data any) {
	app.jsonResponse(c, http.StatusOK, data)
}

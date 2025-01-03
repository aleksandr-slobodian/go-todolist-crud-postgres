package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponseEnvelope struct {
	Data any `json:"data"`
}
func (app *application) jsonResponse(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponseEnvelope{Data: data})
}

func (app *application) jsonOkResponse(c *gin.Context, data any) {
	app.jsonResponse(c, http.StatusOK, data)
}

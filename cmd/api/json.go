package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponseEnvelopeStrict[T any] struct {
	Data T `json:"data"`
}
type jsonResponseEnvelope struct {
	Data any `json:"data"`
}
func (app *application) jsonResponse(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponseEnvelope{Data: data})
}

func (app *application) jsonOkResponse(c *gin.Context, data any) {
	app.jsonResponse(c, http.StatusOK, data)
}

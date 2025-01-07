package main

import (
	"github.com/gin-gonic/gin"
)

type healthCheckResponse struct {
	Status  string `json:"status"` 
	Env     string `json:"env"`     
	Version string `json:"version"` 
}


// healthcheckHandler godoc
//
//	@Summary		Health Check
//	@Description	Returns the health status, environment, and version of the application.
//	@Accept			json
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	jsonResponseEnvelope{data=healthCheckResponse}
//	@Failure		500	{object}	jsonErrorResponseEnvelope
//	@Router			/health [get]
func (app *application) healthCheckHandler(c *gin.Context) {
	app.jsonOkResponse(c, healthCheckResponse{
		Status:  "ok",
		Env:     app.config.env,
		Version: app.config.version,
	})
}

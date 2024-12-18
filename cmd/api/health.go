package main

import (
	"github.com/gin-gonic/gin"
)

type HealthData struct {
	Status  string `json:"status"` 
	Env     string `json:"env"`     
	Version string `json:"version"` 
}

type HealthResponse struct {
	Data HealthData `json:"data"`
}

// healthCheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Provides a health check endpoint to verify the application's status.
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	HealthResponse	"JSON object with health status"
//	@Router			/health [get]
func (app *application) healthCheckHandler(c *gin.Context) {
	app.jsonOkResponse(c, HealthData{
		Status:  "ok",
		Env:     app.config.env,
		Version: app.config.version,
	})
}

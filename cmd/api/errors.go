package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/cmd/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (app *application) errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			if appErr, ok := c.Errors.Last().Err.(*AppError); ok {
					c.JSON(appErr.Code, c.Errors.Last())
			} else {
				c.JSON(http.StatusInternalServerError, c.Errors.Last())
			}
				// TODO: Implement logger
			r := c.Request
			log.Printf("Error ==> method: %s path: %s message: %s", r.Method, r.URL.Path, c.Errors.Last().Error())
		}
	}
}

func (app *application) newError(c *gin.Context, code int, message string, cause ...string) {
	var msg = message
	if len(cause) > 0 && app.config.env == "development" {
		msg = fmt.Sprintf("%s: %s", message, cause[0])
	}
	
	err := &AppError{
		Code:    code,
		Message: msg,
	}
	
	c.Error(err)
	c.Abort()
}

func (app *application) internalServerError(c *gin.Context) {
	app.newError(c, http.StatusInternalServerError, "internal server error")
}

func (app *application) forbiddenResponse(c *gin.Context){
	app.newError(c, http.StatusForbidden, "forbidden")
}

func (app *application) badRequestResponse(c *gin.Context, err error){
	app.newError(c, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(c *gin.Context, err error){
	app.newError(c, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(c *gin.Context){
	app.newError(c, http.StatusNotFound, "not found")
}

func (app *application) notFoundStoreResponse(c *gin.Context, err error){
	switch {
		case errors.Is(err, store.ErrNotFound): 
			app.notFoundResponse(c)
		default: 
			app.internalServerError(c)
		}
}

func (app *application) unauthorizedErrorResponse(c *gin.Context) {	
	app.newError(c, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicErrorResponse(c *gin.Context) {	
	c.Writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	app.newError(c, http.StatusUnauthorized, "unauthorized", "basic error")
}

func (app *application) rateLimitExceededResponse(c *gin.Context, retryAfter string){
	c.Writer.Header().Set("Retry-After", retryAfter)
	app.newError(c, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}

func (app *application) parseValidationError(err error, errType string) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var result string
		for _, fieldError := range validationErrors {
			result += fmt.Sprintf(
				"%s validation for '%s' failed: '%s' (condition: %s)",
				errType,
				fieldError.Field(),
				fieldError.ActualTag(),
				fieldError.Param(),
			)
		}
		return errors.New(result)
	}
	return errors.New("an unknown validation error occurred")
}

func (app *application) shouldBindJSON(c *gin.Context, payload interface{}) error {
	const maxBytes = 1 << 20  // 1 MB

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
	if err := c.ShouldBindJSON(payload); err != nil {
		if err.Error() == "http: request body too large" {
			return errors.New("request body exceeds the maximum allowed size of 1 MB")
		}
		if errors.Is(err, io.EOF) {
			return errors.New("request body is empty")
		}
		return app.parseValidationError(err, "field")
	}

	return nil
}

func (app *application) shouldBindUri(c *gin.Context, payload interface{}) error {
	if err := c.ShouldBindUri(payload); err != nil {
		return app.parseValidationError(err, "param");
	}
	return nil
}
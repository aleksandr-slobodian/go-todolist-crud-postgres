package main

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func numericString(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", fl.Field().String())
	return matched
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("numericString", numericString)
	}
}
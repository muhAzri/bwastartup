package helper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{Message: message, Code: code, Status: status}

	jsonResponse := Response{Meta: meta, Data: data}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}


func IsEmailExistsError(err error) bool {
	var ErrEmailExists = errors.New("email already exists")

	if err == ErrEmailExists {
		return true
	}

	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
			return true
		}
		if strings.Contains(errMsg, "23505") {
			return true
		}
	}

	return false
}

func ErrorResponse(c *gin.Context, message string, statusCode int, errors interface{}) {
	errorMessage := gin.H{"errors": errors}
	response := APIResponse(message, statusCode, "error", errorMessage)
	c.JSON(statusCode, response)
}
package handlers

import (
	"github.com/aerosystems/stat-service/internal/models"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

type BaseHandler struct {
	eventRepo models.EventRepository
}

func NewBaseHandler(
	eventRepo models.EventRepository,
) *BaseHandler {
	return &BaseHandler{
		eventRepo: eventRepo,
	}
}

// Response is the type used for sending JSON around
type Response struct {
	Message string `json:"message,omitempty"`
	Total   int    `json:"total,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// ErrResponse is the type used for sending JSON around
type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// SuccessResponse takes a response status code and arbitrary data and writes a json response to the client in depends on Header Accept
func SuccessResponse(c echo.Context, statusCode int, message string, data any) error {
	payload := Response{
		Message: message,
		Data:    data,
	}
	return c.JSON(statusCode, payload)
}

// ErrorResponse takes a response status code and arbitrary data and writes a json response to the client in depends on Header Accept and APP_ENV environment variable(has two possible values: dev and prod)
// - APP_ENV=dev responds debug info level of error
// - APP_ENV=prod responds just message about error [DEFAULT]
func ErrorResponse(c echo.Context, statusCode, errorCode int, message string, err error) error {
	// TODO: add custom codes for errors
	payload := ErrResponse{
		Code:    errorCode,
		Message: message,
	}

	if strings.ToLower(os.Getenv("APP_ENV")) == "dev" {
		payload.Data = err.Error()
	}

	return c.JSON(statusCode, payload)
}

package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, Response{
		Status:  "error",
		Message: message,
	})
}

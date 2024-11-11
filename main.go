package main

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
)

var task string

type RequestBody struct{
	Message string `json:"message"`
}

func GetHandler(c echo.Context) error{
	return c.String(http.StatusOK, "hello,"+task)
}

func PostHandler(c echo.Context) error {
	var requestBody RequestBody

	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "ошибочка")
	}

	task = requestBody.Message
	return c.String(http.StatusOK, "норм все")
}

func main() {
	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)
	e.Start(":8080")
}


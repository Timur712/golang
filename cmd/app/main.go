package main

import (
	"log"
	"myproject/internal/database"
	"myproject/internal/taskService"
	"myproject/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	repository := taskService.NewTaskRepository()
	service := taskService.NewTaskService(repository)

	e := echo.New()
	handlers.RegisterRoutes(e, service)

	log.Fatal(e.Start(":8080"))
}

package handlers

import (
	"myproject/internal/taskService"
	"myproject/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var taskSvc taskService.TaskService

func RegisterRoutes(e *echo.Echo, service taskService.TaskService) {
	taskSvc = service
	e.GET("/tasks", GetTaskHandler)
	e.POST("/tasks", PostTasksHandler)
	e.PATCH("/tasks/:id", PatchTaskHandler)
	e.DELETE("/tasks/:id", DeleteTaskHandler)
}

func GetTaskHandler(c echo.Context) error {
	tasks, err := taskSvc.GetAllTasks()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Не удалось получить задачи")
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostTasksHandler(c echo.Context) error {
	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Неверный ввод данных")
	}

	newTask, err := taskSvc.CreateTask(task)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Не удалось создать задачу")
	}

	return c.JSON(http.StatusCreated, newTask)
}

func PatchTaskHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Неверный ID")
	}

	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Неверный ввод данных")
	}

	updatedTask, err := taskSvc.UpdateTask(uint(id), task)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Не удалось обновить задачу")
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTaskHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Неверный ID")
	}

	if err := taskSvc.DeleteTask(uint(id)); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Не удалось удалить задачу")
	}

	return utils.SuccessResponse(c, "Задача успешно удалена")
}

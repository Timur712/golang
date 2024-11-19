package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func GetHandler(c echo.Context) error {
	var messages []Message

	if err:= DB.Find(&messages).Error; err != nil{
		return c.JSON(http.StatusBadRequest, Response{
			Status: "error",
			Message: "Couldnt find the message",
		})	
	}
	return c.JSON(http.StatusOK ,&messages)

}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil{
			return c.JSON(http.StatusBadRequest, Response{
				Status: "error",
				Message: "Couldnt add message",
			})
	}

	if err :=DB.Create(&message).Error; err!=nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "error",
			Message: "Couldnt add message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status: "Succes",
		Message: "message was successfully added)))))",
	})

}
func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Bad ID !!!",
		})
	}

	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "неверный ввод",
		})
	}

	if err := DB.Model(&Message{}).Where("id = ?", id).Update("text",updatedMessage.Text).Error; err != nil{
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Couldn't update message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message updated",
	})
}

func DeleteHandler(c echo.Context) error{
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Bad ID !!!",
		})
	}
	if err:= DB.Delete(&Message{},id).Error; err != nil{
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Couldn't delete message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message deleted",
	})
	
}



 func main (){
	e:= echo.New()
	InitDB()
	DB.AutoMigrate(&Message{})

	e.GET("/messages",GetHandler)
	e.POST("/messages",PostHandler)
	e.PATCH("/messages/:id",PatchHandler)
	e.DELETE("/messages/:id",DeleteHandler)
	e.Start(":8080")

 }
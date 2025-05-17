package main

import (
	"net/http"

	"example.com/mygolangproj/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/createEvent", createEvent)
	server.PUT("/updateEvent", updateEvent)
	server.DELETE("/deleteEvent/:id", deleteEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusAccepted, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, "request is invalid")
		return
	}
	event.Save()
	context.JSON(http.StatusAccepted, gin.H{"body": "successfully created the event"})
}

func updateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, "input request is invalid")
		return
	}
	models.UpdateEvent(event.ID, event)
	context.JSON(http.StatusAccepted, gin.H{"body": "updated the event successfully ", "event": event})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	// id := context.Query("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, "id is not provided")
		return
	}
	models.DeleteEvent(id)
	context.JSON(http.StatusAccepted, gin.H{"body": "deleted the event successfully"})
}

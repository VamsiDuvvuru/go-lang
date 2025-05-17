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

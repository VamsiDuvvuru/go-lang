package routes

import (
	"net/http"

	"example.com/mygolangproj/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusAccepted, events)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, "request is invalid")
		return
	}
	event.Save()
	context.JSON(http.StatusAccepted, gin.H{"body": "successfully created the event"})
}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, "input request is invalid")
		return
	}
	models.UpdateEvent(event.ID, event)
	context.JSON(http.StatusAccepted, gin.H{"body": "updated the event successfully ", "event": event})
}

func DeleteEvent(context *gin.Context) {
	id := context.Param("id")
	// id := context.Query("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, "id is not provided")
		return
	}
	models.DeleteEvent(id)
	context.JSON(http.StatusAccepted, gin.H{"body": "deleted the event successfully"})
}

func GetEventsByID(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, "id is not provided")
	}
	var event models.Event = models.GetEventsById(id)
	if event.ID == 0 && event.Name == "" {
		context.JSON(http.StatusNotFound, gin.H{"body": "event not found"})
		return
	}
	context.JSON(http.StatusAccepted, event)
}

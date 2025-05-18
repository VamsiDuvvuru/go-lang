package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.POST("/createEvent", CreateEvent)
	server.PUT("/updateEvent", UpdateEvent)
	server.DELETE("/deleteEvent/:id", DeleteEvent)
	server.GET("/getEvent/:id", GetEventsByID)
}

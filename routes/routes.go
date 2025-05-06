package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/event", createEvents)
	server.GET("/event/:id", getEvent)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)
	server.POST("/signup", signUp)
}

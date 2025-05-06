package main

import (
	"event_planning_go/db"
	"event_planning_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.Run(":8081")
}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later.", "error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not parse request data.","error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	err =  event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not create event.","error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "event created!", "event": event})
}

package routes

import (
	"event_planning_go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later.", "error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvents(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == ""{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err.Error()})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could fetch the event with that event id", "error": err.Error()})
		return
	}
	
	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	event.ID = eventId
	err = event.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event updated!", "event": event})
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err.Error()})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could fetch the event with that event id", "error": err.Error()})
		return
	}
	
	var event models.Event

	event.ID = eventId
	err = event.Delele()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event deleted!"})
}
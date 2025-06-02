package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func CreateEvent(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("CreateEvent")

	event := models.Event{}
	err := c.ShouldBindJSON(&event)
	if err != nil {
		logger.Print("Error: Data binding issue", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	userId := c.GetInt64("userId")
	event.UserId = userId
	//event.DateTime = time.Now()
	id, err := event.Save()
	if err != nil {
		logger.Print("Error: could not create event", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
	event.ID = id
	logger.Print("Event created ok, event", event)
	c.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

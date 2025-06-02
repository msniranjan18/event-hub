package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func UpdateEvent(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("UpdateEvent")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Print("Error: could not parse the event id", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the event id"})
		return
	}

	oldEvent, err := models.GetEventById(eventId)
	if err != nil {
		logger.Print("Error: could not find the event", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not find the event"})
		return
	}

	userId := c.GetInt64("userId")
	if oldEvent.UserId != userId {
		logger.Print("Error: could not find the event", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not Authorized"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		logger.Print("Error: could not find the event", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	updatedEvent.ID = oldEvent.ID
	updatedEvent.DateTime = oldEvent.DateTime
	updatedEvent.UserId = oldEvent.UserId
	err = updatedEvent.Update()
	if err != nil {
		logger.Print("Error: could not find the event", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event"})
		return
	}
	logger.Print("Event Update OK; oldEvent", oldEvent, "updatedEvent", updatedEvent)
	c.JSON(http.StatusAccepted, gin.H{"updated event": updatedEvent})

}

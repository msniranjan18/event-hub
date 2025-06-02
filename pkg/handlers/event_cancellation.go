package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func EventCancellation(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("EventCancellation")

	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Print("Error: could not parse the event id", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the event id"})
		return
	}

	registrations, err := models.GetRegistrationsByUserId(userId)
	if err != nil {
		logger.Print("Error: unable to fetch the registrations", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to fetch the registrations"})
		return
	}
	FOUND := false
	for _, r := range registrations {
		if eventId == r.EventId {
			FOUND = true
			break
		}
	}
	if !FOUND {
		if err != nil {
			logger.Print("Error: unable to fetch the event", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "unable to fetch the event"})
			return
		}
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.Print("Error: unable to fetch the event", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to fetch the event"})
		return
	}

	err = event.Cancel(userId)
	if err != nil {
		logger.Print("Error: could not find the event", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not find the event"})
		return
	}
	logger.Print("OK: successfully cancelled the event", event)
	c.JSON(http.StatusOK, gin.H{"message": "successfully cancelled the event"})
}

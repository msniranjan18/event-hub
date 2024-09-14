package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func GetUsersEventRegistrations(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("GetUsersEventRegistrations")

	userId := c.GetInt64("userId")
	registrations, err := models.GetRegistrationsByUserId(userId)
	if err != nil {
		logger.Print("Error: could not fetch event")
		c.JSON(http.StatusNotFound, gin.H{"message": "could not fetch event"})
		return
	}
	var events []models.Event
	for _, r := range registrations {
		var event models.Event
		event, err = models.GetEventById(r.EventId)
		if err != nil {
			logger.Print("Error: could not fetch event", err.Error())
			c.JSON(http.StatusNotFound, gin.H{"message": "could not fetch event"})
		}
		events = append(events, event)
	}

	logger.Print("OK registered events", events)
	c.JSON(http.StatusOK, gin.H{"events": events})
}

package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func EventRegistration(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("EventRegistration")

	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Print("Error: could not parse the event", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.Print("Error: could not find the event", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not find the event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		logger.Print("Error: could not register", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
		return
	}
	logger.Print("event register OK User ", userId, " eventId ", eventId)
	c.JSON(http.StatusOK, gin.H{"message": "successfully registered for the event"})
}

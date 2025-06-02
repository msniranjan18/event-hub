package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func DeleteEvent(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("DeleteEvent")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Print("Error: could not parse the event id", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.Print("Error: could not fetch event", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	userId := c.GetInt64("userId")
	if event.UserId != userId {
		logger.Print("Error: not Authorized", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not Authorized"})
		return
	}

	err = event.Delete()
	if err != nil {
		logger.Print("Error: could not delete event", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}
	logger.Print("OK: delete event", event)
	c.JSON(http.StatusNoContent, gin.H{"message": "event deleted"})
}

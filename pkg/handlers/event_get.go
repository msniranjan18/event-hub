package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func GetEventbyId(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("GetEventbyId")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Print("Error: could not parse the event id", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.Print("Error: could not fetch event", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "could not fetch event"})
		return
	}
	logger.Print("OK, get event", event)
	c.JSON(http.StatusOK, event)
}

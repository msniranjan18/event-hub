package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func GetEvents(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("GetEvents")

	events, err := models.GetAllEvent()
	if err != nil {
		logger.Print("Error:  unable to fetch the events", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal issue, unable to fetch the events at the moment"})
		return
	}
	logger.Print("OK: List events", events)
	c.JSON(http.StatusOK, gin.H{"events": events})
}

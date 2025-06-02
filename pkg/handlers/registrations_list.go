package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func GetAllRegistrations(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("GetAllRegistrations")

	registrations, err := models.GetAllRegistrations()
	if err != nil {
		logger.Print("Error: unable to fetch the registrations", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal issue, unable to fetch the registrations at the moment"})
		return
	}
	logger.Print("list registrations OK ", registrations)
	c.JSON(http.StatusOK, gin.H{"events": registrations})
}

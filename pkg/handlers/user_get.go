package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func GetUserByEmailID(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("GetUserByEmailID")

	emailID := c.Param("email")

	user, err := models.GetUserByEmailId(emailID)
	if err != nil {
		logger.Print("Error: could not fetch user", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "could not fetch user"})
		return
	}
	logger.Print("user get OK; user:", user)
	c.JSON(http.StatusOK, user)
}

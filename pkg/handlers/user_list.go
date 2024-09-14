package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func GetUsers(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("GetUsers")

	users, err := models.GetAllUsers()
	if err != nil {
		logger.Print("Error: unable to fetch the users", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal issue, unable to fetch the users at the moment"})
		return
	}
	logger.Print("users list OK, users: ", users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

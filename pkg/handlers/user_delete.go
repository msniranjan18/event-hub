package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func DeleteUser(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("DeleteUser")

	emailId := c.Param("email")
	user, err := models.GetUserByEmailId(emailId)
	if err != nil {
		logger.Print("Error: could not fetch user ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
		return
	}

	err = user.Delete()
	if err != nil {
		logger.Print("Error: could not fetch user ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete user"})
		return
	}
	logger.Print("user delete OK user:", user)
	c.JSON(http.StatusNoContent, gin.H{"message": "user deleted"})
}

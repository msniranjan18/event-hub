package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func UpdateUser(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("UpdateUser")

	emailID := c.Param("email")
	oldUser, err := models.GetUserByEmailId(emailID)
	if err != nil {
		logger.Print("Error:could not find the user", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the user"})
		return
	}

	var updatedUser models.User
	updatedUser.UserID = oldUser.UserID
	updatedUser.Email = oldUser.Email
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		logger.Print("Error: data binding issue", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	err = updatedUser.Update()
	if err != nil {
		logger.Print("Error: could not update the user ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the user"})
		return
	}
	logger.Print("successfully updated user")
	c.JSON(http.StatusAccepted, gin.H{"message": "successfully updated user", "New User": updatedUser})

}

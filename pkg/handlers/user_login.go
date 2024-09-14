package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
	"msn.com/event-hub/pkg/utils"
)

func UserLogin(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("UserLogin")

	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		logger.Print("Error: unable to parse the given payload", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse the given payload"})
		return
	}

	userId, err := user.ValidateCredentials()
	if err != nil {
		logger.Print("Error: unable to validate", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	tokenStr, err := utils.GenerateToken(user.Email, userId)
	if err != nil {
		logger.Print("Error: cannot generate the token", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate the token"})
		//return
	}
	logger.Print("login sucsessful, UserId:", userId, " token: ", tokenStr)
	c.JSON(http.StatusOK, gin.H{"message": "login sucsessful", "token": tokenStr})
}

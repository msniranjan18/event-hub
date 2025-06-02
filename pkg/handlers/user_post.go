package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/models"
)

func CreateUser(c *gin.Context) {
	// Retrieve the logger from the context
	loggerC, _ := c.Get("logger")
	logger, _ := loggerC.(*log.Logger)
	logger.Print("CreateUser")

	// Read the raw body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Print("Error: Unable to read body", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read body"})
		return
	}

	// Log or dump the body
	bodyString := string(bodyBytes)
	logger.Printf("Incoming request body: %s", bodyString)

	// Restore the body to the context so it can be read again
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	user := models.User{}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		logger.Print("Error: Data binding error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err})
		return
	}

	id, err := user.Save()
	if err != nil {
		logger.Print("Error: could not create user", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}
	user.UserID = id
	logger.Print("Successfully Registed")
	c.JSON(http.StatusCreated, gin.H{"message": "Successfully Registed", "user": user})
}

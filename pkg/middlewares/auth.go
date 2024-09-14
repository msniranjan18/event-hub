package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/utils"
)

func Authenticate(c *gin.Context) {
	// JWT token verification
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not Authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set("userId", userId)
	c.Next()
}

package middlewares

import (
	"net/http"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	uID, err := utils.ValidateToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set("userID", uID)
	c.Next()

}

package middleware

import (
	"net/http"
	"api-gateway-service/api/token"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")

	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	_, err := token.ExtractClaims(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "invalid token",
		})
		return
	}

	c.Next()
}
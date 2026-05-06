package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const API_KEY = "4p1K3y60L4ng"

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("x-api-key")

		if key == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "API Key required",
			})
			c.Abort()
			return
		}

		if key != API_KEY {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API Key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

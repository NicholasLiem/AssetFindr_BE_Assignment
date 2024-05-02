package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplyJSONMiddleware(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Content-Type") != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Content type must be application/json"})
			c.Abort()
			return
		}
		h(c)
	}
}

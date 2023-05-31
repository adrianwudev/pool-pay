package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	log.Println("authMiddleware")

	headerToken := c.Request.Header.Get("token")
	if headerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
		c.Abort()
		return
	}

	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	log.Println("adminMiddleware")
}

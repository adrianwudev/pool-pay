package middleware

import (
	"log"
	"net/http"
	"pool-pay/internal/auth"

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

	email := auth.GetEmailFromJWT(headerToken)

	log.Printf("original email in middleware: %s\n", email)

	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	log.Println("adminMiddleware")
}

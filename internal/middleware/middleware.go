package middleware

import (
	"errors"
	"log"
	"net/http"
	"pool-pay/internal/auth"
	"pool-pay/internal/util"
	redis_client "pool-pay/redis"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	log.Println("authMiddleware")

	headerToken := c.Request.Header.Get("token")
	if headerToken == "" {
		response := util.NewErrorResponse(errors.New("token is required"))
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	// token is existed in Redis or not
	emailJWT := auth.GetEmailFromJWT(headerToken)
	emailRedis, err := redis_client.GetFromRedis(headerToken)
	if err != nil {
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	log.Printf("emailJWT: %s, emailRedis: %s\n", emailJWT, emailRedis)

	// check if email from JWT is identical with token from Redis
	if emailJWT != emailRedis {
		response := util.NewErrorResponse(errors.New("token is invalid"))
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	log.Println("adminMiddleware")
}

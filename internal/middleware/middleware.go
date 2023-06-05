package middleware

import (
	"log"
	"net/http"
	"pool-pay/internal/auth"
	"pool-pay/internal/constants"
	"pool-pay/internal/util"
	redis_client "pool-pay/redis"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	log.Println("authMiddleware")

	headerToken := c.Request.Header.Get("token")
	if headerToken == "" {
		abortWhenError(util.SetApiError(constants.ERRORCODE_TOKENISREQUIRED), c)
		return
	}
	// token is existed in Redis or not
	emailJWT := auth.GetEmailFromJWT(headerToken)
	emailRedis, err := redis_client.GetFromRedis(headerToken)
	if err != nil {
		if err.(*util.ApiError).Code == constants.ERRORCODE_KEYNOTFOUND {
			err = util.SetApiError(constants.ERRORCODE_TOKENISEXPIRED)
		}
		abortWhenError(err, c)
		return
	}

	log.Printf("emailJWT: %s, emailRedis: %s\n", emailJWT, emailRedis)

	// check if email from JWT is identical with token from Redis
	if emailJWT != emailRedis {
		abortWhenError(util.SetApiError(constants.ERRORCODE_TOKENISINVALID), c)
		return
	}

	// refresh token expired time
	err = redis_client.RefreshExpiredTime(headerToken)
	if err != nil {
		abortWhenError(err, c)
		return
	}

	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	log.Println("adminMiddleware")
}

func abortWhenError(err error, c *gin.Context) {
	response := util.NewErrorResponse(err, util.GetApiError(err).Code)
	c.JSON(http.StatusUnauthorized, response)
	c.Abort()
}

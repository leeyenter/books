package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leeyenter/books/backend/utils"
	"net/http"
)

func GetUserIdFromHeader(c *gin.Context) (string, error) {
	token := c.GetHeader("auth")
	if token == "" {
		return "", fmt.Errorf("no auth header found")
	}
	return parseJWT(token)
}

func Middleware(c *gin.Context) {
	tokenIp, err := GetUserIdFromHeader(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	requestIp := utils.GetRemoteIp(c.Request)
	if tokenIp != requestIp {
		fmt.Println(tokenIp, requestIp)
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	c.Next()
}

package middlewares

import (
	"net/http"

	"example.com/car-rental/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	context.Set("UserId", userId)
	context.Next()
}
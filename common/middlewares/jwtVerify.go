package middlewares

import (
	"github.com/Singh555/mycms/common/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(http.StatusGone, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}

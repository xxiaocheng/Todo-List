package utils

import (
	"net/http"
	"time"
	"todoList/serializers"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errorCode int
		var data interface{}
		h := serializers.AuthorizationHeaderRequest{}
		if err := c.ShouldBindHeader(&h); err != nil {
			errorCode = serializers.ErrorAuth
		} else {
			claims, err := ParseJwtToken(h.Authorization)
			if err != nil {
				errorCode = serializers.ErrorAuth
			} else {
				if claims.VerifyExpiresAt(time.Now().Unix(), true) {
					errorCode = serializers.ErrorAuthCheckTokenTimeout
				}
			}
		}
		if errorCode != serializers.Success {
			c.JSON(http.StatusForbidden, serializers.BuildResponse(errorCode, data))
			c.Abort()
			return
		}
		c.Next()
	}
}

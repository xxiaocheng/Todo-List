package middlewares

import (
	"net/http"
	"time"
	"todoList/models"
	"todoList/serializers"
	"todoList/utils/jwt"

	"github.com/gin-gonic/gin"
)

func updateContextUserModel(c *gin.Context, username string) {
	user, _ := models.FindOneUserByUsername(username)
	c.Set("username", username)
	c.Set("userModel", user)
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errCode int
		var data interface{}
		errCode = serializers.Success
		h := serializers.AuthorizationHeaderRequest{}
		appG := serializers.Gin{C: c}
		if err := c.ShouldBindHeader(&h); err != nil {
			errCode = serializers.ErrorAuth
		} else {
			h.StripBearerPrefix()
			claims, err := jwt.ParseJwtToken(h.Authorization)
			if err != nil {
				errCode = serializers.ErrorAuth
			} else {
				if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
					errCode = serializers.ErrorAuthCheckTokenTimeout
				} else {
					updateContextUserModel(c, (*claims)["username"].(string))
				}
			}
		}
		if errCode != serializers.Success {
			appG.Response(http.StatusUnauthorized, errCode, data)
			c.Abort()
			return
		}
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE, PATCH")
		c.Set("content-type", "application/json")
		if c.Request.Method=="OPTIONS"{
			c.AbortWithStatus(http.StatusNoContent)
			c.Header("Allow","POST, GET, HEAD, OPTIONS")
		}
		c.Next()
	}
}

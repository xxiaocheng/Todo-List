package v1

import (
	"net/http"
	"todoList/serializers"
	"todoList/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserLogin(c *gin.Context) {
	loginRequest := serializers.LoginRequest{}
	appG := serializers.Gin{C: c}

	if err := c.ShouldBind(&loginRequest); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err := validate.Struct(&loginRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	ok, err := loginRequest.CheckAuth()
	if err != nil {
		appG.Response(http.StatusInternalServerError, serializers.ErrorAuthCheckTokenFail, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusUnauthorized, serializers.ErrorAuth, nil)
		return
	}
	token := jwt.GenJwtToken(loginRequest.Username)
	c.JSON(http.StatusOK, token)
}

func UserRegister(c *gin.Context) {
	registerRequest := serializers.RegisterUserRequest{}
	appG := serializers.Gin{C: c}

	if err := c.ShouldBind(&registerRequest); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err := validate.Struct(&registerRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	_, err = registerRequest.Register()
	if err != nil {
		appG.Response(http.StatusUnprocessableEntity, serializers.ErrorPutUserFail, nil)
		return
	}
	appG.Response(http.StatusOK, serializers.Success, nil)
}

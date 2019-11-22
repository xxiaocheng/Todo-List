package v1

import (
	"net/http"
	"todoList/serializers"
	"todoList/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Get Jwt
// @Tags Auth
// @Accept multipart/form-data
// @Produce  json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} serializers.JwtResponse "OK"
// @Failure 401 {object} serializers.CommonResponse "FAIL"
// @Router /auth/token [post]
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
		appG.Response(http.StatusUnauthorized, serializers.ErrorAuthCheckTokenFail, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusUnauthorized, serializers.ErrorAuth, nil)
		return
	}
	token := jwt.GenJwtToken(loginRequest.Username)
	c.JSON(http.StatusOK, token)
}

// @Summary User Register
// @Tags Auth
// @Accept multipart/form-data
// @Produce  json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param email formData string true "email"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /auth/register [post]
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

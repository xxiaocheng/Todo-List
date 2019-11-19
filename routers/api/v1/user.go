package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"todoList/models"
	"todoList/serializers"
)

func ChangeUserPassword(c *gin.Context) {
	appG := serializers.Gin{C: c}
	r := serializers.ModifyPasswordRequest{}
	if err := c.ShouldBind(&r); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err := validate.Struct(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	userModel := (c.MustGet("userModel")).(models.User)
	err = userModel.ChangePassword(r.Password)
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorChangePasswordFail, nil)
		return
	}
	appG.Response(http.StatusOK, serializers.Success, nil)
}

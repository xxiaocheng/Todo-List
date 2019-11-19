package v1

import (
	"net/http"
	"todoList/models"
	"todoList/serializers"
	"todoList/utils/hashID"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetGroups(c *gin.Context) {
	pageDataRequest := serializers.PageDataRequest{}
	appG := serializers.Gin{C: c}
	if err := c.ShouldBindQuery(&pageDataRequest); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	pageDataRequest.ValidDefault()
	userModel := (c.MustGet("userModel")).(models.User)
	groups, count, err := userModel.PagedGroups(pageDataRequest.Offset, pageDataRequest.Limit)
	if err != nil {
		appG.Response(http.StatusInternalServerError, serializers.ErrorGetGroupFail, nil)
		return
	}
	cr := serializers.PagedDataCommonResponse{Offset: pageDataRequest.Offset,
		Limit: pageDataRequest.Limit,
		Count: count,
		Data:  serializers.SerializerGroupsFromModel(groups),
	}
	appG.Response(http.StatusOK, serializers.Success, cr)
}

func DeleteOneGroup(c *gin.Context) {
	appG := serializers.Gin{C: c}
	groupHashID := c.Param("group")
	userModel := (c.MustGet("userModel")).(models.User)
	groupID := hashID.DecodeHashToID(groupHashID)
	err := userModel.DeleteOneGroup(groupID)
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorDeleteGroupFail, nil)
		return
	} else {
		appG.Response(http.StatusOK, serializers.Success, nil)
		return
	}
}

func ModifyOneGroupName(c *gin.Context) {
	appG := serializers.Gin{C: c}
	// get param
	r := serializers.ModifyGroupRequest{}
	err := c.ShouldBind(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err = validate.Struct(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	groupHashID := c.Param("group")
	userModel := (c.MustGet("userModel")).(models.User)
	groupID := hashID.DecodeHashToID(groupHashID)
	err = userModel.ModifyGroupName(groupID, r.NewGroupName)
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorEditGroupFail, nil)
		return
	} else {
		appG.Response(http.StatusOK, serializers.Success, nil)
		return
	}
}

func CreateOneGroup(c *gin.Context) {
	appG := serializers.Gin{C: c}
	// get param
	r := serializers.CreateGroupRequest{}
	err := c.ShouldBind(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err = validate.Struct(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	userModel := (c.MustGet("userModel")).(models.User)
	group, err := userModel.CreateOneGroup(r.GroupName)
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorEditGroupFail, nil)
		return
	} else {
		appG.Response(http.StatusCreated, serializers.Success, serializers.GroupSerializer(group))
		return
	}
}

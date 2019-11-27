package v1

import (
	"net/http"
	"todoList/models"
	"todoList/serializers"
	"todoList/utils/hashID"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Paged all groups
// @Tags Group
// @Accept multipart/form-data
// @Produce  json
// @Security ApiKeyAuth
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /group/ [get]
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

// @Summary Delete one group by groupID
// @Tags Group
// @Accept multipart/form-data
// @Produce  json
// @Security ApiKeyAuth
// @Param group path string true "group hashID"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /group/{group} [delete]
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

// @Summary Rename group
// @Tags Group
// @Accept multipart/form-data
// @Produce  json
// @Security ApiKeyAuth
// @Param group path string true "group hashID"
// @Param new_group_name formData string true "new group name"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /group/{group} [patch]
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

// @Summary Create one  group
// @Tags Group
// @Accept multipart/form-data
// @Produce  json
// @Security ApiKeyAuth
// @Param group_name formData string true "group name"
// @Success 201 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /group/ [post]
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

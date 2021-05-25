package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAuth(c *gin.Context) {
	if res := haveAuth(c, "getAllAuth"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authModel := model.Auth{}

	authJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Title string `json:"title" form:"title"`
		}
	}{}

	if c.ShouldBind(&authJson) == nil {
		authJson.Offset = (authJson.Offset - 1) * authJson.Limit
		res := authModel.GetAllAuth(authJson.Offset, authJson.Limit, authJson.Where.Title)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetParentAuth(c *gin.Context) {
	if res := haveAuth(c, "getAllAuth"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authModel := model.Auth{}

	authJson := struct {
		Type int `json:"type" form:"type"`
	}{}

	if err := c.ShouldBind(&authJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}


	res := authModel.GetParentAuth(authJson.Type)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return

}

func AddAuth(c *gin.Context) {
	if res := haveAuth(c, "addAuth"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	var authJson model.Auth

	if err := c.ShouldBind(&authJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	authMap := helper.Struct2Map(authJson)
	if res, err := authValidate.ValidateMap(authMap, "add"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	authJson.Target = "_self"
	if authJson.Type != 1 {
		authJson.Href = ""
	}
	res := authModel.AddAuth(authJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteAuth(c *gin.Context) {
	if res := haveAuth(c, "deleteAuth"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	authIDJson := struct {
		Aid int `json:"aid" form:"aid"`
	}{}

	if err := c.ShouldBind(&authIDJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	authIDMap := helper.Struct2Map(authIDJson)
	if res, err := authValidate.ValidateMap(authIDMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}
	res := authModel.DeleteAuth(authIDJson.Aid)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return

}

func UpdateAuth(c *gin.Context) {
	if res := haveAuth(c, "updateAuth"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	var authJson model.Auth

	if err := c.ShouldBind(&authJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	authMap := helper.Struct2Map(authJson)
	if res, err := authValidate.ValidateMap(authMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := authModel.UpdateAuth(authJson.Aid, authJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetAuthByID(c *gin.Context) {
	if res := haveAuth(c, "getAllAuth"); res != common.Authed { //getAllUser怎么改？
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	var authJson model.Auth

	if err := c.ShouldBind(&authJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	authMap := helper.Struct2Map(authJson)
	if res, err := authValidate.ValidateMap(authMap, "find"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := authModel.GetAuthByID(authJson.Aid)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

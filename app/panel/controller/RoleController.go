package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRole(c *gin.Context) { //??
	roleModel := model.Role{}

	roleJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Name string `json:"name" form:"name"`
			Desc string `json:"desc" form:"desc"`
		}
	}{}

	if c.ShouldBind(&roleJson) == nil {
		roleJson.Offset = (roleJson.Offset - 1) * roleJson.Limit
		res := roleModel.GetAllRole(roleJson.Offset, roleJson.Limit, roleJson.Where.Name, roleJson.Where.Desc)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

func GetRoleByID(c *gin.Context) { //jun
	roleValidate := validate.RoleValidate
	roleModel := model.Role{}

	var roleJson model.Role

	if err := c.ShouldBind(&roleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	roleMap := helper.Struct2Map(roleJson)
	if res, err := roleValidate.ValidateMap(roleMap, "find"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := roleModel.GetRoleByID(roleJson.Rid)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return

}

func AddRole(c *gin.Context) { //jun
	roleValidate := validate.RoleValidate
	roleModel := model.Role{}

	var roleJson model.Role
	if err := c.ShouldBind(&roleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	roleMap := helper.Struct2Map(roleJson)
	if res, err := roleValidate.ValidateMap(roleMap, "add"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	//userJson.Password = common.GetMd5(userJson.Password)
	res := roleModel.AddRole(roleJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteRole(c *gin.Context) {
	roleValidate := validate.RoleValidate
	roleModel := model.Role{}

	roleIDJson := struct {
		Rid int `json:"rid" form:"rid"`
	}{}

	if err := c.ShouldBind(&roleIDJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	roleIDMap := helper.Struct2Map(roleIDJson)
	if res, err := roleValidate.ValidateMap(roleIDMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := roleModel.DeleteRole(roleIDJson.Rid)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateRole(c *gin.Context) {
	roleValidate := validate.RoleValidate
	roleModel := model.Role{}

	var roleJson model.Role

	if err := c.ShouldBind(&roleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	roleMap := helper.Struct2Map(roleJson)
	if res, err := roleValidate.ValidateMap(roleMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := roleModel.UpdateRole(roleJson.Rid, roleJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

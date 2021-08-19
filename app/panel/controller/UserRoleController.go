package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserRoles(c *gin.Context) {
	userRoleValidate := validate.UserRoleValidate
	userRoleModel := model.UserRole{}

	userRolesJson := struct {
		UserID int    `json:"user_id" form:"user_id"`
		Rids   string `json:"rids" form:"rids"`
	}{}

	if err := c.ShouldBind(&userRolesJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userRolesMap := helper.Struct2Map(userRolesJson)
	if res, err := userRoleValidate.ValidateMap(userRolesMap, "addGroup"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	var rids []int
	_ = json.Unmarshal([]byte((userRolesJson.Rids)), &rids)
	fmt.Println(rids)
	for _, rid := range rids {
		res := userRoleModel.AddUserRole(model.UserRole{UserID: userRolesJson.UserID, Rid: rid})
		if res.Status != constants.CodeSuccess {
			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, "编号为"+string(rune(rid))+"的角色添加失败", res.Data))
			return
		}
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "添加成功", true))
	return
}

func DeleteUserRoles(c *gin.Context) {
	userRoleValidate := validate.UserRoleValidate
	userRoleModel := model.UserRole{}

	userRolesJson := struct {
		UserID int    `json:"user_id" form:"user_id"`
		Rids   string `json:"rids" form:"rids"`
	}{}

	if err := c.ShouldBind(&userRolesJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	userRolesMap := helper.Struct2Map(userRolesJson)
	if res, err := userRoleValidate.ValidateMap(userRolesMap, "deleteGroup"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	var rids []int
	_ = json.Unmarshal([]byte((userRolesJson.Rids)), &rids)
	fmt.Println(rids)
	for _, rid := range rids {
		res := userRoleModel.DeleteUserRole(model.UserRole{UserID: userRolesJson.UserID, Rid: rid})
		if res.Status != constants.CodeSuccess {
			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, "编号为"+string(rune(rid))+"的权限删除失败", res.Data))
			return
		}
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "删除成功", true))
	return

}

func GetUserRolesList(c *gin.Context) {
	userRoleValidate := validate.UserRoleValidate
	roleModel := model.Role{}

	roleJson := struct {
		UserID int `json:"user_id" form:"user_id"`
	}{}

	if err := c.ShouldBind(&roleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	roleMap := helper.Struct2Map(roleJson)
	if res, err := userRoleValidate.ValidateMap(roleMap, "getUserRole"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	allRoles := roleModel.GetRoleNoRules()

	res := roleModel.GetUserRole(roleJson.UserID)
	roles := res.Data.([]model.Role)
	var val []int
	for _, role := range roles {
		val = append(val, role.Rid)
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, map[string]interface{}{
		"allRoles": allRoles.Data,
		"values":   val,
	}))
	return
}

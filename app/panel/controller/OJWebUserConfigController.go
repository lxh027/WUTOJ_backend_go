package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AddOJWebUserConfig 添加用户配置
func AddOJWebUserConfig(c *gin.Context) {
	ojWebUserConfigValidate := validate.OJWebUserConfigValidate
	ojWebUserConfigModel := model.OJWebUserConfig{}
	var ojWebUserConfigJSON model.OJWebUserConfig
	if err := c.ShouldBind(&ojWebUserConfigJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	ojWebUserConfigMap := helper.Struct2Map(ojWebUserConfigJSON)
	if res, err := ojWebUserConfigValidate.ValidateMap(ojWebUserConfigMap, "add"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	res := ojWebUserConfigModel.AddOJWebUserConfig(ojWebUserConfigJSON)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

//DeleteOJWebUserConfig 删除用户配置
func DeleteOJWebUserConfig(c *gin.Context) {
	ojWebUserConfigModel := model.OJWebUserConfig{}
	ojWebUserConfigValidate := validate.OJWebUserConfigValidate
	var ojWebUserConfigJSON model.OJWebUserConfig
	if err := c.ShouldBind(&ojWebUserConfigJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	ojWebUserConfigMap := helper.Struct2Map(ojWebUserConfigJSON)
	if res, err := ojWebUserConfigValidate.ValidateMap(ojWebUserConfigMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	res := ojWebUserConfigModel.DeleteOJWebUserConfig(ojWebUserConfigJSON.ID)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

//GetOJWebUserConfigByID 由ID获取OJ用户配置
func GetOJWebUserConfigByID(c *gin.Context) {
	ojWebUserConfigModel := model.OJWebUserConfig{}
	ojWebUserConfigValidate := validate.OJWebUserConfigValidate
	var ojWebUserConfigJSON model.OJWebUserConfig

	if err := c.ShouldBind(&ojWebUserConfigJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	ojWebUserConfigMap := helper.Struct2Map(ojWebUserConfigJSON)
	if res, err := ojWebUserConfigValidate.ValidateMap(ojWebUserConfigMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := ojWebUserConfigModel.GetOJWebUserConfigByID(ojWebUserConfigJSON.ID)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	ojWebUserConfigRes := res.Data.(model.OJWebUserConfig)
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "get success", ojWebUserConfigRes))
	return
}

//UpdateOJWebUserConfig 修改用户OJ配置
func UpdateOJWebUserConfig(c *gin.Context) {
	ojWebUserConfigModel := model.OJWebUserConfig{}
	ojWebUserConfigValidate := validate.OJWebUserConfigValidate
	var ojWebUserConfigJSON model.OJWebUserConfig
	if err := c.ShouldBind(&ojWebUserConfigJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(ojWebUserConfigJSON)
	if res, err := ojWebUserConfigValidate.ValidateMap(contestMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := ojWebUserConfigModel.UpdateOJWebUserConfig(ojWebUserConfigJSON.ID, ojWebUserConfigJSON)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

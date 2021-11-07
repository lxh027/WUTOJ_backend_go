package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

//DeleteOJWebUserData 删除某用户在特定OJ的所有数据
func DeleteOJWebUserData(c *gin.Context) {
	ojWebUserDataModel := model.OJWebUserData{}
	ojWebDataValidate := validate.OJWebDataValidate
	var ojWebUserDataJSON model.OJWebUserData
	if err := c.ShouldBind(&ojWebUserDataJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	ojWebUserDataMap := helper.Struct2Map(ojWebUserDataJSON)
	if res, err := ojWebDataValidate.ValidateMap(ojWebUserDataMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	res := ojWebUserDataModel.DeleteOJWebUserData(ojWebUserDataJSON.UserID, ojWebUserDataJSON.OJName)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

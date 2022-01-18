package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

//DeleteOJWebUserData 删除用户在特定OJ的某条数据
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
	res := ojWebUserDataModel.DeleteOJWebUserData(ojWebUserDataJSON.UserID, ojWebUserDataJSON.OJName, ojWebUserDataJSON.SubmitTime)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

//GetAllOJWebUserData 获取所有用户OJ做题信息
func GetAllOJWebUserData(c *gin.Context) {
	dataModel := model.OJWebUserData{}

	dataJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			UserID int    `json:"user_id" form:"user_id"`
			OJName string `json:"oj_name" form:"oj_name"`
		}
	}{}

	if c.ShouldBind(&dataJson) == nil {
		dataJson.Offset = (dataJson.Offset - 1) * dataJson.Limit
		res := dataModel.GetAllOJWebUserData(dataJson.Offset, dataJson.Limit, dataJson.Where.UserID, dataJson.Where.OJName)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

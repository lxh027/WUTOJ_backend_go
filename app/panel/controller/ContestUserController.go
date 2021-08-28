package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllContestUsers 获取所有参赛选手
func GetAllContestUsers(c *gin.Context) {
	contestUserModel := model.ContestUser{}

	contestuserJSON := struct {
		ContestID int `json:"contest_id" form:"contest_id"`
	}{}

	if c.ShouldBind(&contestuserJSON) == nil {
		res := contestUserModel.GetAllContestUsersByID(contestuserJSON.ContestID)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

//自建

//GetAllContestUsers 获取所有参赛选手
func GetAllContestUsers(c *gin.Context) {
	contestUserModel := model.ContestUser{}

	contestuserJSON := struct {
		ContestID int `json:"contest_id" form:"contest_id" uri:"contest_id"`
		UserID    int `json:"user_id" form:"user_id"`
		Status    int `json:"status" form:"status"`
	}{}

	if c.ShouldBind(&contestuserJSON) == nil {
		// contestuserJSON.Offset = (noticeJson.Offset - 1) * noticeJson.Limit
		res := contestUserModel.GetAllContestUsersByID(contestuserJSON.ContestID)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

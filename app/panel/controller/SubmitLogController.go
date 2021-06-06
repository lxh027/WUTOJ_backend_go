package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUserSubmitStatus(c *gin.Context) {
	userLogModel := model.UserSubmitLog{}

	userLogJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Nick string `json:"nick" form:"nick"`
		}
	}{}

	if c.ShouldBind(&userLogJson) == nil {
		userLogJson.Offset = (userLogJson.Offset - 1) * userLogJson.Limit
		res := userLogModel.GetAllUserSubmitStatus(userLogJson.Offset, userLogJson.Limit, userLogJson.Where.Nick)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

func GetUserSubmitStatusByTime(c *gin.Context) {
	userLogModel := model.UserSubmitLog{}

	userLogJson := struct {
		UserId []int `json:"user_id" form:"user_id"`
		StartTime  string `json:"start_time" form:"start_time"`
		EndTime string `json:"end_time" form:"end_time"`
	}{}

	if c.ShouldBind(&userLogJson) == nil {
		fmt.Fprintf(gin.DefaultWriter, "userLogJson: %v", userLogJson)
		res := userLogModel.GetUserSubmitStatusByTime(userLogJson.UserId, userLogJson.StartTime, userLogJson.EndTime)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}
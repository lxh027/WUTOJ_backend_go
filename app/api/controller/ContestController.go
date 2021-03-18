package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
)

func GetAllContest(c *gin.Context) {
	contestModel := model.Contest{}
	res := contestModel.GetAllContest()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetContest(c *gin.Context) {
	ContestID := c.Param("contest_id")
	fmt.Println(ContestID)
	contestModel := model.Contest{}
	var res helper.ReturnType
	res = contestModel.GetContestById(ContestID)
	if res.Status != common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		res = contestModel.GetContestByName(ContestID)
		if res.Status != common.CodeError {
			c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
			return
		} else {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "查找失败", ""))
			return
		}

	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func JoinContest(c *gin.Context) {
	var contestUserModel = model.ContestUser{}
	var contestUserJson struct {
		model.ContestUser
	}
	if err := c.ShouldBind(&contestUserJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
	res := contestUserModel.AddContestUser(contestUserJson.ContestUser)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

func GetContestStatus(c *gin.Context) {
	var contestModel = model.Contest{}
	var ContestID struct {
		ID int `form:contest_id`
	}
	if err := c.ShouldBind(&ContestID); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
	res := contestModel.GetContestStatus(ContestID.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserContest(c *gin.Context) {
	var queryInfo struct {
		UserID int `form:"user_id"`
	}
	contestUserModel := model.ContestUser{}
	if err := c.ShouldBind(&queryInfo); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", ""))
		return
	}
	res := contestUserModel.GetUserContest(queryInfo.UserID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

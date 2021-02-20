package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
)

func isInt(contestID string) bool {
	if contestID[0] > '0' && contestID[0] < '9' {
		return true
	} else {
		return false
	}
}

func GetAllContest(c *gin.Context) {
	contestModel := model.Contest{}
	res := contestModel.GetAllContest()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetContest(c *gin.Context) {
	ContestID := c.Param("id")
	contestModel := model.Contest{}
	var res helper.ReturnType
	res = contestModel.GetContestById(ContestID)
	if res.Status != common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	} else {
		res = contestModel.GetContestByName(ContestID)
		if res.Status != common.CodeError {
			c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		} else {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "查找失败", ""))
		}
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func AddContest(c *gin.Context) {
	var contestModel = model.Contest{}
	var contestJson struct {
		model.Contest
	}
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	res := contestModel.AddContest(contestJson.Contest)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

func DeleteContest(c *gin.Context) {
	var contestModel = model.Contest{}
	var ContestID struct {
		ID int `form:contest_id`
	}
	if err := c.ShouldBind(&ContestID); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
	res := contestModel.DeleteContest(ContestID.ID)
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

}

func UpdateContest(c *gin.Context) {
	var contestModel = model.Contest{}
	var contestJson struct {
		model.Contest
	}
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	res := contestModel.UpdateContest(contestJson.ContestID, contestJson.Contest)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

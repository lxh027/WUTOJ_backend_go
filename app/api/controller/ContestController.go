package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func isInt(contestID string) bool {
	if contestID[0] > '0' && contestID[0] < '9' {
		return true
	} else {
		return false
	}
}

func GetAll(c *gin.Context) {
	contestModel := model.Contest{}
	res := contestModel.GetAllContest()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetContest(c *gin.Context) {
	ContestID := c.Param("id")
	contestModel := model.Contest{}
	var res helper.ReturnType
	if isInt(ContestID) {
		res = contestModel.GetContestById(ContestID)
	} else {
		res = contestModel.GetContestByName(ContestID)
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func Add(c *gin.Context) {
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

package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAllProblems(c *gin.Context) {

	problemModel := model.Problem{}
	res := problemModel.GetAllProblems()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}

func GetProblemByID(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}
	contestModel := model.Contest{}

	var problemJson model.Problem

	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	// TODO: need remove, temprory workaround
	contestJson := contestModel.GetContestByProblemId(problemMap["problem_id"].(int))
	if contestJson.Status == common.CodeSuccess {
		now := time.Now()
		contest := contestJson.Data.(model.Contest)
		if now.Before(contest.BeginTime) || contest.EndTime.Before(now) {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "比赛未开始", 0))
			return
		}
	}

	res := problemModel.GetProblemByID(int(problemJson.ProblemID))
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func SearchProblem(c *gin.Context) {
	problemJson := struct {
		Param string `uri:"param" json:"param"`
	}{}

	problemJson.Param = c.Param("param")
	problemModel := model.Problem{}
	contestModel := model.Contest{}

	if err := c.ShouldBind(&problemJson); err == nil {
		res := problemModel.SearchProblem(problemJson.Param)

		// TODO: need remove, temprory workaround
		problem_id, _ := strconv.Atoi(problemJson.Param)
		contestJson := contestModel.GetContestByProblemId(problem_id)
		if contestJson.Status == common.CodeSuccess {
			now := time.Now()
			contest := contestJson.Data.(model.Contest)
			if now.Before(contest.BeginTime) || contest.EndTime.Before(now) {
				c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "比赛未开始", 0))
				return
			}
		}
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
}

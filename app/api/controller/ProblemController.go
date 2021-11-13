package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"net/http"
	"strconv"
	"time"
	"log"
	"github.com/gin-gonic/gin"
)

func GetAllProblems(c *gin.Context) {

	problemModel := model.Problem{}
	problemJson := struct {
		PageNumber int `json:"page_number" form:"page_number"`
	}{}
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "页码参数错误", err.Error()))
		return
	}
	res := problemModel.GetAllProblems((problemJson.PageNumber-1)*constants.PageLimit, constants.PageLimit)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}

func GetProblemByID(c *gin.Context) {
	userIDRaw := GetUserIdFromSession(c)

	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}
	contestModel := model.Contest{}
	contestUserModel := model.ContestUser{}

	problemJson := struct {
		ProblemID int `json:"problem_id"`
	}{}
	if problemID, err := strconv.Atoi(c.Param("problem_id")); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "参数错误", err.Error()))
		return
	} else {
		problemJson.ProblemID = problemID
	}
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.GetProblemByID(int(problemJson.ProblemID))
	log.Printf("\n\n%v\n\n", res.Data.(map[string]interface{}))
	if res.Status != constants.CodeSuccess || res.Data.(map[string]interface{})["problem"].(model.Problem).Public == constants.ProblemPublic {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	// judge if problem in contests
	contestsBeginTime := contestModel.GetContestsByProblemID(
		problemJson.ProblemID,
		[]string{"contest.contest_id", "begin_time"},
	)

	if contestsBeginTime.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(contestsBeginTime.Status, contestsBeginTime.Msg, contestsBeginTime.Msg))
	}
	for _, contest := range contestsBeginTime.Data.([]model.Contest) {
		userID := int(userIDRaw)
		if participation := contestUserModel.CheckUserContest(userID, contest.ContestID); participation.Status == constants.CodeSuccess {
			format := "2006-01-02 15:04:05"
			now, _ := time.Parse(format, time.Now().Format(format))
			beginTime, _, _, err := getContestTime(uint(contest.ContestID))
			if err == nil && now.Unix() >= beginTime.Unix() {
				//res := problemModel.GetProblemByID(int(problemJson.ProblemID))
				c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
				return
			}
		}
	}
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "没有权限", 0))
	return

	/*
		//
		contestJson := contestModel.GetContestByProblemId(problemMap["problem_id"].(int))
		if contestJson.Status == constants.CodeSuccess {

			contest := contestJson.Data.(model.Contest)
			userID := int(userIDRaw)
			contestID := contest.ContestID
			if participation := contestUserModel.CheckUserContest(userID,contestID); participation.Status != constants.CodeSuccess{
				c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "尚未参赛，请参赛", 0))
				return
			}

			format := "2006-01-02 15:04:05"
			now, _ := time.Parse(format, time.Now().Format(format))
			if now.Before(contest.BeginTime) || contest.EndTime.Before(now) {
				c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "比赛未开始", 0))
				return
			}
		}*/

}

func SearchProblem(c *gin.Context) {
	userJson := checkLogin(c)
	if userJson.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "未登录", 0))
		return
	}
	userIDRaw := userJson.Data.(uint)

	problemJson := struct {
		Param string `uri:"param" json:"param"`
	}{}

	problemJson.Param = c.Param("param")
	problemModel := model.Problem{}
	contestModel := model.Contest{}
	contestUserModel := model.ContestUser{}

	if err := c.ShouldBind(&problemJson); err == nil {
		res := problemModel.SearchProblem(problemJson.Param)

		// TODO: need remove, temprory workaround
		problemId, _ := strconv.Atoi(problemJson.Param)
		contestsBeginTime := contestModel.GetContestsByProblemID(
			problemId,
			[]string{"contest.contest_id", "begin_time"},
		)

		if contestsBeginTime.Status != constants.CodeSuccess {
			c.JSON(http.StatusOK, helper.ApiReturn(contestsBeginTime.Status, contestsBeginTime.Msg, contestsBeginTime.Msg))
		}

		for _, contest := range contestsBeginTime.Data.([]model.Contest) {
			userID := int(userIDRaw)
			if participation := contestUserModel.CheckUserContest(userID, contest.ContestID); participation.Status != constants.CodeSuccess {
				c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "No participation for contest found", nil))
				return
			}
			format := "2006-01-02 15:04:05"
			now, _ := time.Parse(format, time.Now().Format(format))
			if now.Before(contest.BeginTime) {
				c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "比赛未开始", 0))
				return
			}
		}

		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
}

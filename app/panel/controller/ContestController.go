package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllContest(c *gin.Context) {
	contestModel := model.Contest{}

	contestJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			ContestName string    `json:"contest_name" form:"contest_name"`
			BeginTime   time.Time `json:"begin_time" form:"begin_time"`
		}
	}{}

	if c.ShouldBind(&contestJson) == nil {
		contestJson.Offset = (contestJson.Offset - 1) * contestJson.Limit
		res := contestModel.GetAllContest(contestJson.Offset, contestJson.Limit, contestJson.Where.ContestName, contestJson.Where.BeginTime)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

func GetContestByID(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}
	contestProblemModel := model.ContestProblem{}

	var contestJson model.Contest

	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.FindContestByID(contestJson.ContestID)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	contestRes := res.Data.(model.Contest)
	problemsRes := contestProblemModel.GetContestProblems(contestJson.ContestID)
	problemStr, _ := json.Marshal(problemsRes.Data)
	contestRes.Problems = string(problemStr)
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "get success", contestRes))
	return
}

func AddContest(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}
	problemModel := model.Problem{}
	contestProblemModel := model.ContestProblem{}
	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "add"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	var problems []int
	if err := json.Unmarshal([]byte(contestJson.Problems), &problems); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	for _, problemId := range problems {
		problemModel.ChangeProblemPublicStatus(problemId, 0)
	}

	res := contestModel.AddContest(contestJson)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	contestID := res.Data.(int)
	for _, id := range problems {
		_ = contestProblemModel.AddContestProblem(model.ContestProblem{ContestID: contestID, ProblemID: id})
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteContest(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}
	contestProblemModel := model.ContestProblem{}
	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	_ = contestProblemModel.DeleteContestProblem(contestJson.ContestID)
	res := contestModel.DeleteContest(contestJson.ContestID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateContest(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}
	problemModel := model.Problem{}
	contestProblemModel := model.ContestProblem{}
	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	var problems []int
	if err := json.Unmarshal([]byte(contestJson.Problems), &problems); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	for _, problemId := range problems {
		problemModel.ChangeProblemPublicStatus(problemId, 0)
	}

	res := contestModel.UpdateContest(contestJson.ContestID, contestJson)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	_ = contestProblemModel.DeleteContestProblem(contestJson.ContestID)
	for _, id := range problems {
		_ = contestProblemModel.AddContestProblem(model.ContestProblem{ContestID: contestJson.ContestID, ProblemID: id})
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeContestStatus(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.ChangeContestStatus(contestJson.ContestID, contestJson.Status)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ClearContestRedis(c *gin.Context) {
	contestValidate := validate.ContestValidate

	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	if err := database.DeleteFromRedis(redis_key.ContestRank(int(contestJson.ContestID))); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "刷新排行榜失败", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "刷新排行榜成功", 0))
	return
}

func SetOuterBoard(c *gin.Context) {
	contestJson := struct {
		ContestID 	int		`json:"contest_id" form:"contest_id"`
		Time 		int 	`json:"time" form:"time"`
	}{}

	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	_ = database.DeleteFromRedis(redis_key.OuterInfo)
	_ = database.DeleteFromRedis(redis_key.OuterSubmits)
	_ = database.DeleteFromRedis(redis_key.OuterTeams)
	_ = database.DeleteFromRedis(redis_key.OuterID)
	if err := database.PutToRedis(redis_key.OuterID, contestJson.ContestID, contestJson.Time); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "redis error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "开放成功", nil))
}
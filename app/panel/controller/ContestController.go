package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"OnlineJudge/core/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddContest(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}
	problemModel := model.Problem{}

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
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteContest(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

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

	res := contestModel.DeleteContest(contestJson.ContestID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateContest(c *gin.Context) {
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}
	problemModel := model.Problem{}

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


	if err := db.DeleteFromRedis("contest_rank" + strconv.Itoa(int(contestJson.ContestID))); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "刷新排行榜失败", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "刷新排行榜成功", 0))
	return
}

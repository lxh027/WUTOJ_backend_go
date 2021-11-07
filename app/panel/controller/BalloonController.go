package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetContestBalloon 获取气球
func GetContestBalloon(c *gin.Context) {
	contestModel := model.Contest{}
	SubmitModel := model.Submit{}

	contestIDJson := struct {
		ContestID uint `json:"contest_id" form:"contest_id"`
	}{}

	if err := c.ShouldBind(&contestIDJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestRes := contestModel.GetContestById(contestIDJson.ContestID)
	if contestRes.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, contestRes.Msg, contestRes.Data))
		return
	}

	submitRes := SubmitModel.GetContestACSubmitsWithExtraName(contestIDJson.ContestID)
	if submitRes.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, submitRes.Msg, submitRes.Data))
		return
	}

	contest := contestRes.Data.(model.Contest)
	submits := submitRes.Data.([]model.SubmitBalloon)

	colors := make([]string, 0)
	problems := make([]uint, 0)
	_ = json.Unmarshal([]byte(contest.Colors), &colors)
	_ = json.Unmarshal([]byte(contest.Problems), &problems)
	fmt.Fprintf(gin.DefaultWriter, "colors: %v\n problems: %v\n", colors, problems)

	problemIDMap := make(map[uint]int)
	for index, problemID := range problems {
		problemIDMap[problemID] = index
	}

	type balloon struct {
		ID        uint   `json:"id"`
		UserID    uint   `json:"user_id"`
		Nick      string `json:"nick"`
		Realname  string `json:"realname" form:"realname"`
		ProblemID int    `json:"problem_id"`
		Color     string `json:"color"`
		IsSent    bool   `json:"is_sent"`
	}

	var balloons []balloon
	balloonMap := make(map[string]bool)

	for _, submit := range submits {
		newBalloon := balloon{
			UserID:    submit.UserID,
			ID:        submit.ID,
			Nick:      submit.Nick,
			ProblemID: problemIDMap[submit.ProblemID] + 1,
			Color:     colors[problemIDMap[submit.ProblemID]],
			Realname:  submit.Realname,
		}
		submitIdentity := strconv.Itoa(int(contestIDJson.ContestID)) + strconv.Itoa(newBalloon.ProblemID) + " " + strconv.Itoa(int(newBalloon.UserID))
		if _, ok := balloonMap[submitIdentity]; ok {
			continue
		}
		balloonMap[submitIdentity] = true
		value, err := database.SIsNumberOfRedisSet(redis_key.Balloon(int(contestIDJson.ContestID)), submitIdentity)
		if err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "Redis错误", err.Error()))
			return
		}
		newBalloon.IsSent = value
		balloons = append(balloons, newBalloon)
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "获取成功", balloons))
}

func SentBalloon(c *gin.Context) {
	IDJson := struct {
		ContestID uint `json:"contest_id" form:"contest_id"`
		ProblemID int  `json:"problem_id" form:"problem_id"`
		UserID    uint `json:"user_id" form:"user_id"`
	}{}

	if err := c.ShouldBind(&IDJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	submitIdentity := strconv.Itoa(int(IDJson.ContestID)) + strconv.Itoa(IDJson.ProblemID) + " " + strconv.Itoa(int(IDJson.UserID))

	if err := database.SAddToRedisSet(redis_key.Balloon(int(IDJson.ContestID)), submitIdentity); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "设置失败", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "设置成功", nil))
}

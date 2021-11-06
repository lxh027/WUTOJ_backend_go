package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"OnlineJudge/core/database"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetContestUsr(c *gin.Context)  {
	userModel := model.User{}
	submitJson := struct {
		ContestID 	int 	`uri:"contest_id"`
	}{}

	if c.ShouldBindUri(&submitJson) != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", nil))
		return
	}

	res1 := userModel.GetContestUser(submitJson.ContestID, 0)
	res2 := userModel.GetContestUser(submitJson.ContestID, 1)

	if res1.Status != constants.CodeSuccess || res2.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "get data error", nil))
		return
	}

	formalUsers := res1.Data.([]model.User)
	starUsers := res2.Data.([]model.User)

	type item map[string]string
	reData := make(map[string]item)

	// formal teams
	for _, user := range formalUsers {
		userItem := item{
			"members": user.Realname,
			"school": user.School,
			"team": user.Nick,
			"type": "type1",
		}
		reData[strconv.Itoa(int(user.UserID))] = userItem
	}

	// star teams
	for _, user := range starUsers {
		userItem := item{
			"members": user.Realname,
			"school": user.School,
			"team": user.Nick,
			"type": "type2",
		}
		reData[strconv.Itoa(int(user.UserID))] = userItem
	}
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "get teams success", reData))
}

func GetContestSubmit(c *gin.Context)  {
	submitModel := model.Submit{}
	contestModel := model.Contest{}
	submitJson := struct {
		ContestID 	int 	`uri:"contest_id"`
	}{}

	if c.ShouldBindUri(&submitJson) != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", nil))
		return
	}

	// try get from redis
	if d, err := database.GetFromRedis("outer_submits"); err == nil && d != nil {
		var reData [][]interface{}
		bytes, _ := redis.Bytes(d, err)
		_ = json.Unmarshal(bytes, &reData)
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "get rank(redis) success", reData))
		return
	}
	// get contest begin time
	res := contestModel.GetContestById(submitJson.ContestID)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	beginTime := res.Data.(model.Contest).BeginTime
	problemsStr := res.Data.(model.Contest).Problems

	var problems []uint

	if json.Unmarshal([]byte(problemsStr), &problems) != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "parse problems error", nil))
		return
	}

	// get submits
	res = submitModel.GetContestSubmits(uint(submitJson.ContestID))

	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	//res
	submits := res.Data.([]model.Submit)

	var reData []interface{}
	for _, submit := range submits {
		data := make([]interface{}, 0)
		var status string
		switch submit.Status {
		case "ac":
			status = "AC"
		case "judging":
			status = "NEW"
		default:
			status = "NO"
		}
		var problemID string
		for index, problem := range problems {
			if problem == submit.ProblemID {
				problemID = string(rune('A' + index))
			}
		}
		data = append(data, submit.UserID, problemID, submit.SubmitTime.Unix()-beginTime.Unix(), status)
		reData = append(reData, data)
	}
	reDataStr, _ := json.Marshal(reData)
	_ = database.PutToRedis("outer_submits", reDataStr, 3600)
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "get rank success", reData))
}


package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"OnlineJudge/core/database"
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetContestInfo(c *gin.Context) {
	var contestID int
	if d, err := database.GetFromRedis("outer_id"); err == nil && d != nil {
		contestID, _ = redis.Int(d, err)
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "no board is opened", nil))
		return
	}

	userInfo, err1 := getContestUsr(contestID)
	submitsInfo, err2 := getContestSubmit(contestID)
	basicInfo, err3 := getContestBasicInfo(contestID)

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "get data error", nil))
		return
	}
	reData := map[string]interface{} {
		"teams": userInfo,
		"submits": submitsInfo,
		"basic": basicInfo,
	}
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "get data success", reData))
}

func getContestUsr(contestID int) (interface{}, error)  {
	type item map[string]string
	userModel := model.User{}

	// try get from redis
	if d, err := database.GetFromRedis("outer_teams"); err == nil && d != nil {
		var reData map[string]item
		bytes, _ := redis.Bytes(d, err)
		_ = json.Unmarshal(bytes, &reData)
		return reData, nil
	}

	res1 := userModel.GetContestUser(contestID, 0)
	res2 := userModel.GetContestUser(contestID, 1)

	if res1.Status != constants.CodeSuccess || res2.Status != constants.CodeSuccess {
		return nil, errors.New("get data error")
	}

	formalUsers := res1.Data.([]model.User)
	starUsers := res2.Data.([]model.User)

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
	reDataStr, _ := json.Marshal(reData)
	_ = database.PutToRedis("outer_teams", reDataStr, 3600)
	return reData, nil
}

func getContestSubmit(contestID int) (interface{}, error)  {
	submitModel := model.Submit{}
	contestModel := model.Contest{}

	// try get from redis
	if d, err := database.GetFromRedis("outer_submits"); err == nil && d != nil {
		var reData [][]interface{}
		bytes, _ := redis.Bytes(d, err)
		_ = json.Unmarshal(bytes, &reData)
		return reData, nil
	}
	// get contest begin time
	res := contestModel.GetContestById(contestID)
	if res.Status != constants.CodeSuccess {
		return nil, errors.New(res.Msg)
	}
	beginTime := res.Data.(model.Contest).BeginTime
	problemsStr := res.Data.(model.Contest).Problems

	var problems []uint

	if json.Unmarshal([]byte(problemsStr), &problems) != nil {
		return nil, errors.New("parse problems error")
	}

	// get submits
	res = submitModel.GetContestSubmits(uint(contestID))

	if res.Status != constants.CodeSuccess {
		return nil, errors.New(res.Msg)
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
	_ = database.PutToRedis("outer_submits", reDataStr, 30)
	return reData, nil
}

func getContestBasicInfo(contestID int) (interface{}, error) {
	contestModel := model.Contest{}

	// try get from redis
	if d, err := database.GetFromRedis("outer_info"); err == nil && d != nil {
		var reData map[string]interface{}
		bytes, _ := redis.Bytes(d, err)
		_ = json.Unmarshal(bytes, &reData)
		return reData, nil
	}

	res := contestModel.GetContestById(contestID)
	if res.Status != constants.CodeSuccess {
		return nil, errors.New(res.Msg)
	}
	contest := res.Data.(model.Contest)
	problemsStr := contest.Problems

	var problems []uint

	if json.Unmarshal([]byte(problemsStr), &problems) != nil {
		return nil, errors.New("parse problems error")
	}

	reData := map[string]interface{} {
		"end_time": contest.EndTime,
		"problem_num": len(problems),
		"title": contest.ContestName,
	}
	reDataStr, _ := json.Marshal(reData)
	_ = database.PutToRedis("outer_info", reDataStr, 3600)
	return reData, nil
}

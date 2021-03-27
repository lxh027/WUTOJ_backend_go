package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/db_server"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type problem struct {
	SuccessTime 	int64 	`json:"success_time"`
	Times 			uint 	`json:"times"`
}
type user struct {
	UserID 	uint 	`json:"user_id"`
	Nick 	string 	`json:"nick"`
	Penalty int64 	`json:"penalty"`
	ACNum 	uint 	`json:"ac_num"`
	ProblemID 	map[uint]problem 	`json:"problem_id"`
}
type userSort []user
func (a userSort) Len() int           { return len(a) }
func (a userSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a userSort) Less(i, j int) bool {
	if a[i].ACNum != a[j].ACNum {
		return a[i].ACNum < a[j].ACNum
	} else {
		return a[i].Penalty < a[j].Penalty
	}
}

func GetUserRank(c *gin.Context)  {
	rankValidate := validate.RankValidate

	rankJson := struct {
		ContestID 	uint 	`uri:"contest_id"`
	}{}

	if err := c.ShouldBindUri(&rankJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	rankMap := helper.Struct2Map(rankJson)
	if res, err := rankValidate.ValidateMap(rankMap, "getContestRank"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	beginTime, endTime, frozenTime, err := getContestTime(rankJson.ContestID)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), nil))
		return
	}

	now := time.Now()

	var begin, end time.Time
	begin = beginTime
	if now.Unix() < beginTime.Unix() {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "比赛未开始", nil))
		return
	} else if now.Unix() < frozenTime.Unix() {
		end = now
	} else if now.Unix() < endTime.Unix() {
		end = frozenTime
	} else {
		end = endTime
	}

	// try get from redis
	if rank, err := db_server.ZGetAllFromRedis("contest_rank"+strconv.Itoa(int(rankJson.ContestID))); err == nil {
		userIDRank, _ := redis.Strings(rank, err)
		var rankBoard []interface{}
		for _, userID := range userIDRank {
			// get each user_id
			var item user
			itemStr, err := redis.String(db_server.GetFromRedis("contest_rank"+strconv.Itoa(int(rankJson.ContestID))+"user_id"+userID))
			if err != nil {
				getRankFromDB(c, rankJson.ContestID, begin, end, now)
				return
			}
			_ = json.Unmarshal([]byte(itemStr), &item)
			rankBoard = append(rankBoard, item)
		}
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取排行榜成功", rankBoard))
		return
	}
	// get from DB
	getRankFromDB(c, rankJson.ContestID, begin, end, now)
}

// TODO get rank from DB
func getRankFromDB(c *gin.Context, contestID uint, beginTime, endTime, now time.Time) {
	submitModel := model.Submit{}

	res := submitModel.GetContestSubmitsByTime(contestID, beginTime, endTime)
	if res.Status != common.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	var users []user
	userIndexMap := make(map[uint]int)
	userProblemMap := make(map[string]bool)
	submits := res.Data.([]model.Submit)

	for _, submit := range submits {
		var index int
		if i, ok := userIndexMap[submit.UserID]; !ok {
			users = append(users, user{UserID: submit.UserID, Nick: submit.Nick, Penalty: 0, ACNum: 0, ProblemID: make(map[uint]problem)})
			index = len(users)-1
		} else {
			index = i
		}

		userIndex := strconv.Itoa(int(submit.UserID))+"."+strconv.Itoa(int(submit.ProblemID))
		if ac, ok := userProblemMap[userIndex]; !ok || !ac {
			// not AC yet
			if _, ok = users[index].ProblemID[submit.ProblemID]; !ok {
				users[index].ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: 0}
			}
			userProblem := users[index].ProblemID[submit.ProblemID]
			if submit.Status == "AC" {
				userProblemMap[userIndex] = true
				users[index].ProblemID[submit.ProblemID] = problem{SuccessTime: now.Unix(), Times: userProblem.Times+1}
				users[index].ACNum++
				for _, problem := range users[index].ProblemID {
					users[index].Penalty += int64(problem.Times*20*60)+problem.SuccessTime
				}
			} else if submit.Status != "CE" {
				userProblemMap[userIndex] = false
				users[index].ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: userProblem.Times+1}
			}
		}
	}

	sort.Sort(userSort(users))
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取排行榜成功", users))

	for _, user := range users {
		itemStr, _ := json.Marshal(user)
		_ = db_server.PutToRedis("contest_rank"+strconv.Itoa(int(contestID))+"user_id"+strconv.Itoa(int(user.UserID)), itemStr, 3600)
		score := fmt.Sprintf("%03d.%d", user.ACNum, user.Penalty)
		_ = db_server.ZAddToRedis("contest_rank"+strconv.Itoa(int(contestID)), score, user.UserID)
	}
}


package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
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
	SuccessTime int64 `json:"success_time"`
	Times       uint  `json:"times"`
}
type user struct {
	UserID    uint             `json:"user_id"`
	Nick      string           `json:"nick"`
	Penalty   int64            `json:"penalty"`
	ACNum     uint             `json:"ac_num"`
	ProblemID map[uint]problem `json:"problem_id"`
}
type userSort []user

func (a userSort) Len() int      { return len(a) }
func (a userSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a userSort) Less(i, j int) bool {
	if a[i].ACNum != a[j].ACNum {
		return a[i].ACNum > a[j].ACNum
	} else {
		return a[i].Penalty < a[j].Penalty
	}
}

func GetUserRank(c *gin.Context) {
	rankValidate := validate.RankValidate

	rankJson := struct {
		ContestID uint `uri:"contest_id"`
	}{}

	if err := c.ShouldBindUri(&rankJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	rankMap := helper.Struct2Map(rankJson)
	if res, err := rankValidate.ValidateMap(rankMap, "getContestRank"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	beginTime, endTime, frozenTime, err := getContestTime(rankJson.ContestID)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), nil))
		return
	}

	format := "2006-01-02 15:04:05"
	now, _ := time.Parse(format, time.Now().Format(format))

	var begin, end time.Time
	begin = beginTime
	fmt.Println(now.String(), beginTime.String(), frozenTime.String(), endTime.String())
	if now.Unix() < beginTime.Unix() {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "比赛未开始", nil))
		return
	} else if now.Unix() < frozenTime.Unix() {
		end = now
	} else if now.Unix() < endTime.Unix() {
		end = frozenTime
	} else {
		end = endTime
	}

	// try get from redis
	if rank, err := database.ZGetAllFromRedis(redis_key.ContestRank(int(rankJson.ContestID))); err == nil {
		userIDRank, _ := redis.Strings(rank, err)
		if len(userIDRank) == 0 {
			getRankFromDB(c, rankJson.ContestID, begin, end, now)
			return
		}
		var rankBoard []interface{}
		for _, userID := range userIDRank {
			// get each user_id
			var item user
			itemStr, err := redis.String(database.GetFromRedis(redis_key.ContestRankUser(int(rankJson.ContestID), userID)))
			if err != nil {
				getRankFromDB(c, rankJson.ContestID, begin, end, now)
				return
			}
			_ = json.Unmarshal([]byte(itemStr), &item)
			rankBoard = append(rankBoard, item)
		}
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "获取排行榜成功", rankBoard))
		return
	}
	// get from DB
	getRankFromDB(c, rankJson.ContestID, begin, end, now)
}

// TODO get rank from DB
func getRankFromDB(c *gin.Context, contestID uint, beginTime, endTime, now time.Time) {
	fmt.Printf("Get From DB between %v and %v\n", beginTime, endTime)
	submitModel := model.Submit{}

	res := submitModel.GetContestSubmitsByTime(contestID, beginTime, endTime)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	fmt.Println(res)
	var users []user
	userIndexMap := make(map[uint]int)
	userProblemMap := make(map[string]bool)
	submits := res.Data.([]model.Submit)

	for _, submit := range submits {
		var index int
		if i, ok := userIndexMap[submit.UserID]; !ok {
			users = append(users, user{UserID: submit.UserID, Nick: submit.Nick, Penalty: 0, ACNum: 0, ProblemID: make(map[uint]problem)})
			index = len(users) - 1
			userIndexMap[submit.UserID] = index
		} else {
			index = i
		}

		userIndex := strconv.Itoa(int(submit.UserID)) + "." + strconv.Itoa(int(submit.ProblemID))
		if ac, ok := userProblemMap[userIndex]; !ok || !ac {
			// not AC yet
			if _, ok = users[index].ProblemID[submit.ProblemID]; !ok {
				users[index].ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: 0}
			}
			userProblem := users[index].ProblemID[submit.ProblemID]
			if submit.Status == "AC" {
				userProblemMap[userIndex] = true
				users[index].ProblemID[submit.ProblemID] = problem{SuccessTime: submit.SubmitTime.Unix() - beginTime.Unix(), Times: userProblem.Times + 1}
				users[index].ACNum++
				users[index].Penalty += int64(userProblem.Times*20*60) + users[index].ProblemID[submit.ProblemID].SuccessTime
			} else if submit.Status != "CE" && submit.Status != "UE"{
				userProblemMap[userIndex] = false
				users[index].ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: userProblem.Times + 1}
			}
		}
	}

	sort.Sort(userSort(users))
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "获取排行榜成功", users))

	for _, user := range users {
		itemStr, _ := json.Marshal(user)
		_ = database.PutToRedis(redis_key.ContestRankUser(int(contestID), strconv.Itoa(int(user.UserID))), itemStr, 3600)
		score := -int64(user.ACNum) * 1000000000 + user.Penalty
		_ = database.ZAddToRedis(redis_key.ContestRank(int(contestID)), score, user.UserID)
	}
}

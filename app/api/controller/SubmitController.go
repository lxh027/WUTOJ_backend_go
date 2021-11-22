package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/config"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"OnlineJudge/core/judger"
	"encoding/json"
	"fmt"
	_ "io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	//TODO: auth participation and contest time
	
	problemModel := model.Problem{}
	contestModel := model.Contest{}
	contestUserModel := model.ContestUser{}

	submitModel := model.Submit{}
	submitValidate := validate.SubmitValidate
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "用户未登录", ""))
		return
	}

	format := "2006-01-02 15:04:05"
	now, _ := time.Parse(format, time.Now().Format(format))
	interval := config.GetWutOjConfig()["interval_time"].(int)
	redisStr := redis_key.UserLastSubmit(int(userID.(uint)))
	if value, err := database.GetFromRedis(redisStr); err == nil {
		defaultFormat := "2006-01-02 15:04:05 +0000 UTC"
		lastStr, _ := redis.String(value, err)
		last, _ := time.Parse(defaultFormat, lastStr)
		fmt.Printf("now: %v, last: %v\n", now, last)

		if now.Unix()-last.Unix() <= int64(interval) {
			c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "交题间隔过快，请五秒后再试", ""))
			return
		}
	}
	_ = database.PutToRedis(redisStr, now, 3600)

	var submitJson struct {
		Language   string `json:"language"`
		SourceCode string `json:"source_code"`
		ProblemID  uint   `json:"problem_id"`
		ContestID  uint   `json:"contest_id"`
	}

	if err := c.ShouldBindJSON(&submitJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据绑定模型错误", err.Error()))
		return
	}

	submitMap := helper.Struct2Map(submitJson)

	if res, err := submitValidate.ValidateMap(submitMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), ""))
		return
	}

	if helper.LanguageID(submitJson.Language) == -1 {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "不支持的语言类型", nil))
		return
	}

	// judge if problem is private and in contests
	ok := false
	res := problemModel.GetProblemByID(int(submitJson.ProblemID))
	if res.Status != constants.CodeSuccess || res.Data.(map[string]interface{})["problem"].(model.Problem).Public == constants.ProblemPublic {
		ok = true
	} else {
		contestsBeginTime := contestModel.GetContestsByProblemID(
			int(submitJson.ProblemID),
			[]string{"contest.contest_id", "begin_time"},
		)
		if contestsBeginTime.Status == constants.CodeSuccess {
			for _, contest := range contestsBeginTime.Data.([]model.Contest) {
				if participation := contestUserModel.CheckUserContest(int(userID.(uint)), contest.ContestID); participation.Status == constants.CodeSuccess {
					format := "2006-01-02 15:04:05"
					now, _ := time.Parse(format, time.Now().Format(format))
					beginTime, _, _, err := getContestTime(uint(contest.ContestID))
					if err == nil && now.Unix() >= beginTime.Unix() {
						ok = true
						break
					}
				}
			}
		}
	}
	if !ok {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "暂时无法提交", nil))
		return
	}
	



	newSubmit := model.Submit{
		UserID:     userID.(uint),
		Nick:       GetUserNickFromSession(c),
		Language:   helper.LanguageID(submitJson.Language),
		SourceCode: submitJson.SourceCode,
		ProblemID:  submitJson.ProblemID,
		ContestID:  submitJson.ContestID,
		Status:     "Judging",
		SubmitTime: now,
	}

	resp := submitModel.AddSubmit(&newSubmit)

	go func(submit model.Submit) {
		judge(submit)
	}(newSubmit)

	c.JSON(http.StatusOK, helper.ApiReturn(resp.Status, resp.Msg, resp.Data))
	return

}

func judge(submit model.Submit) {
	langConfig := config.GetLangConfigs()[submit.Language]
	submitData := judger.SubmitData{
		Id:           uint64(submit.ID),
		Pid:          uint64(submit.ProblemID),
		Language:     langConfig.Lang,
		Code:         submit.SourceCode,
		BuildScript:  langConfig.BuildSh,
		RunnerConfig: langConfig.RunnerConfig,
	}

	callback := func(id uint64, result judger.JudgeResult) {
		// Put Result To DB
		// TODO if set contest, update redis rank info
		if result.IsFinished {
			data := map[string]interface{}{
				"status": result.Status,
				"time":   result.Time,
				"memory": result.Memory,
				"msg":    result.Msg,
			}
			submitModel := model.Submit{}
			submitModel.UpdateStatusAfterSubmit(int(id), data)
			if submit.ContestID != 0 {
				// set to redis
				beginTime, _, frozenTime, err := getContestTime(submit.ContestID)
				if err != nil {
					return
				}
				format := "2006-01-02 15:04:05"
				now, _ := time.Parse(format, time.Now().Format(format))
				fmt.Printf("Submit Time: %v, BeginTime: %v, FrozenTime: %v\n", now, beginTime, frozenTime)
				if now.Unix() < beginTime.Unix() || now.Unix() > frozenTime.Unix() {
					return
				}
				user := user{UserID: submit.UserID, Nick: submit.Nick, Penalty: 0, ACNum: 0, ProblemID: make(map[uint]problem)}
				if itemStr, err := redis.String(database.GetFromRedis(redis_key.ContestRankUser(int(submit.ContestID), strconv.Itoa(int(submit.UserID))))); err == nil {
					_ = json.Unmarshal([]byte(itemStr), &user)
				}
				if _, ok := user.ProblemID[submit.ProblemID]; !ok {
					user.ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: 0}
				}
				userProblem := user.ProblemID[submit.ProblemID]
				log.Println(result)
				if user.ProblemID[submit.ProblemID].SuccessTime == 0 {
					if result.Status == "AC" {
						user.ProblemID[submit.ProblemID] = problem{SuccessTime: submit.SubmitTime.Unix() - beginTime.Unix(), Times: userProblem.Times + 1}
						user.ACNum++
						user.Penalty += int64(userProblem.Times*20*60) + user.ProblemID[submit.ProblemID].SuccessTime
					} else if result.Status != "CE" && result.Status != "UE" {
						user.ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: userProblem.Times + 1}
					}

					itemStr, _ := json.Marshal(user)
					_ = database.PutToRedis(redis_key.ContestRankUser(int(submit.ContestID), strconv.Itoa(int(user.UserID))), itemStr, 3600)
					score := -int64(user.ACNum)*1000000000 + user.Penalty
					_ = database.ZAddToRedis(redis_key.ContestRank(int(submit.ContestID)), score, user.UserID)

				}
			}
		}
	}

	instance := judger.GetInstance()

	go instance.Submit(submitData, callback)
}

func GetSubmitInfo(c *gin.Context) {
	submitModel := model.Submit{}

	submitJson := struct {
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	ConfigMap := config.GetWutOjConfig()
	Limit := ConfigMap["page_limit"].(int)

	UserID := GetUserIdFromSession(c)

	if c.ShouldBind(&submitJson) == nil {
		res := submitModel.GetAllSubmit(Limit*(submitJson.PageNumber-1), Limit, UserID)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", false))

}

func GetProblemSubmitInfo(c *gin.Context) {
	submitModel := model.Submit{}
	submitJson := model.Submit{}

	if c.ShouldBindQuery(&submitJson) != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	submitJson.UserID = GetUserIdFromSession(c)
	res := submitModel.GetProblemSubmit(submitJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetUserContestSubmitInfo(c *gin.Context) {
	submitValidate := validate.SubmitValidate
	submitModel := model.Submit{}

	submitJson := struct {
		ContestID  uint `json:"contest_id" form:"contest_id" uri:"contest_id"`
		PageNumber int  `json:"page_number" form:"page_number" uri:"page_number"`
		UserID     uint
	}{}

	if c.ShouldBindQuery(&submitJson) != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	submitJson.UserID = GetUserIdFromSession(c)

	submitMap := helper.Struct2Map(submitJson)
	if res, err := submitValidate.ValidateMap(submitMap, "contest_log"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := submitModel.GetContestSubmitByUser(submitJson.UserID, submitJson.ContestID, submitJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetSubmitByID(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	submitJson := model.Submit{}
	submitModel := model.Submit{}
	submitValidate := validate.SubmitValidate

	if c.ShouldBindQuery(&submitJson) != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	submitMap := helper.Struct2Map(submitJson)
	if res, err := submitValidate.ValidateMap(submitMap, "find"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	submitJson.UserID = GetUserIdFromSession(c)

	res = submitModel.GetSubmitByID(submitJson.ID, submitJson.UserID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

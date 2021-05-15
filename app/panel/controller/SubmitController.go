package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/db_server"
	"OnlineJudge/judger"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		return a[i].ACNum < a[j].ACNum
	} else {
		return a[i].Penalty < a[j].Penalty
	}
}

func GetAllSubmit(c *gin.Context) {
	if res := haveAuth(c, "getAllSubmit"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitModel := model.Submit{}

	submitJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			UserID        string    `json:"user_id" form:"user_id"`
			ProblemID     string    `json:"problem_id" form:"problem_id"`
			ContestID     string    `json:"contest_id" form:"contest_id"`
			Language      string    `json:"language" form:"language"`
			Status        string    `json:"status" form:"status"`
			MinSubmitTime time.Time `json:"min_submit_time" form:"min_submit_time"`
			MaxSubmitTime time.Time `json:"max_submit_time" form:"max_submit_time"`
		}
	}{}
	var err error
	if err = c.ShouldBind(&submitJson); err == nil {
		submitJson.Offset = (submitJson.Offset - 1) * submitJson.Limit
		whereData := map[string]string{
			"user_id": submitJson.Where.UserID, "problem_id": submitJson.Where.ProblemID,
			"contest_id": submitJson.Where.ContestID, "language": submitJson.Where.Language,
			"status": submitJson.Where.Status,
		}
		res := submitModel.GetAllSubmit(submitJson.Offset, submitJson.Limit, whereData, submitJson.Where.MinSubmitTime, submitJson.Where.MaxSubmitTime)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), false))
	return
}

func GetSubmitByID(c *gin.Context) {
	if res := haveAuth(c, "getAllSubmit"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitValidate := validate.SubmitValidate
	submitModel := model.Submit{}

	var submitJson model.Submit

	if err := c.ShouldBind(&submitJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	submitMap := helper.Struct2Map(submitJson)
	if res, err := submitValidate.ValidateMap(submitMap, "find"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := submitModel.FindSubmitByID(submitJson.ID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func RejudgeGroupSubmits(c *gin.Context) {
	if res := haveAuth(c, "rejudge"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitModel := model.Submit{}

	submitJson := struct {
		UserID        string    `json:"user_id" form:"user_id"`
		ProblemID     string    `json:"problem_id" form:"problem_id"`
		ContestID     string    `json:"contest_id" form:"contest_id"`
		Language      string    `json:"language" form:"language"`
		Status        string    `json:"status" form:"status"`
		MinSubmitTime time.Time `json:"min_submit_time" form:"min_submit_time"`
		MaxSubmitTime time.Time `json:"max_submit_time" form:"max_submit_time"`
	}{}

	if c.ShouldBind(&submitJson) == nil {
		whereData := map[string]string{
			"user_id": submitJson.UserID, "problem_id": submitJson.ProblemID,
			"contest_id": submitJson.ContestID, "language": submitJson.ContestID,
			"status": submitJson.Status,
		}
		res := submitModel.GetSubmitGroup(whereData, submitJson.MinSubmitTime, submitJson.MaxSubmitTime)
		submits := res.Data.([]model.Submit)
		for _, submit := range submits {
			go func(submitData model.Submit) {
				rejudge(submitData)
			}(submit)
		}
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeSuccess, "发送重测成功", true))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func RejudgeSubmitByID(c *gin.Context) {
	if res := haveAuth(c, "rejudge"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitValidate := validate.SubmitValidate
	submitModel := model.Submit{}

	var submitJson model.Submit

	if err := c.ShouldBind(&submitJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	submitMap := helper.Struct2Map(submitJson)
	if res, err := submitValidate.ValidateMap(submitMap, "rejudge"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := submitModel.FindSubmitByID(submitJson.ID)
	submit := res.Data.(model.Submit)
	go func(submitData model.Submit) {
		rejudge(submitData)
	}(submit)
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeSuccess, "发送重测成功", true))
	return
}

func rejudge(submit model.Submit) {
	submitData := judger.SubmitData{
		Id:           uint64(submit.ID),
		Pid:          uint64(submit.ProblemID),
		Language:     helper.LanguageType(submit.Language),
		Code:         submit.SourceCode,
		BuildScript:  "",
		RootfsConfig: nil,
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
				if now.Unix() < beginTime.Unix() || now.Unix() > frozenTime.Unix() {
					return
				}
				user := user{UserID: submit.UserID, Nick: submit.Nick, Penalty: 0, ACNum: 0, ProblemID: make(map[uint]problem)}
				if itemStr, err := redis.String(db_server.GetFromRedis("contest_rank" + strconv.Itoa(int(submit.ContestID)) + "user_id" + strconv.Itoa(int(submit.UserID)))); err == nil {
					_ = json.Unmarshal([]byte(itemStr), &user)
				}
				if _, ok := user.ProblemID[submit.ProblemID]; !ok {
					user.ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: 0}
				}
				userProblem := user.ProblemID[submit.ProblemID]
				log.Println(result)
				if result.Status == "AC" {
					user.ProblemID[submit.ProblemID] = problem{SuccessTime: submit.SubmitTime.Unix()-beginTime.Unix(), Times: userProblem.Times + 1}
					user.ACNum++
					for _, problem := range user.ProblemID {
						user.Penalty += int64(problem.Times*20*60) + problem.SuccessTime
					}
				} else if result.Status != "CE" {
					user.ProblemID[submit.ProblemID] = problem{SuccessTime: 0, Times: userProblem.Times + 1}
				}

				itemStr, _ := json.Marshal(user)
				_ = db_server.PutToRedis("contest_rank"+strconv.Itoa(int(submit.ContestID))+"user_id"+strconv.Itoa(int(user.UserID)), itemStr, 3600)
				score := -int64(user.ACNum) * 1000000000 + user.Penalty
				_ = db_server.ZAddToRedis("contest_rank"+strconv.Itoa(int(submit.ContestID)), score, user.UserID)
			}
		}
	}

	instance := judger.GetInstance()

	go instance.Submit(submitData, callback)
}

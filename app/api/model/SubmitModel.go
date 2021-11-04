package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"

	"github.com/gin-gonic/gin"
)

type Submit struct {
	ID         uint      `json:"id" form:"id"`
	UserID     uint      `json:"user_id" form:"user_id"`
	Nick       string    `json:"nick" form:"nick"`
	ProblemID  uint      `json:"problem_id" form:"problem_id" uri:"problem_id"`
	ContestID  uint      `json:"contest_id" form:"contest_id" uri:"contest_id"`
	SourceCode string    `json:"source_code" form:"source_code"`
	Language   int       `json:"language" form:"language"`
	Status     string    `json:"status" form:"status"`
	Time       int64     `json:"time" form:"time"`
	Memory     uint      `json:"memory" form:"memory"`
	Msg        string    `json:"msg" form:"msg"`
	SubmitTime time.Time `json:"submit_time" form:"submit_time"`
}

var submitFields []string = []string{"id", "user_id", "nick", "problem_id", "contest_id", "source_code", "language", "status", "time", "memory", "msg", "submit_time"}

func (Submit) TableName() string {
	return "submit"
}

func (model *Submit) GetReturnData(submits []Submit) []map[string]interface{} {
	var submitsData []map[string]interface{}

	for _, submit := range submits {
		submitsData = append(submitsData, map[string]interface{}{
			"id":          submit.ID,
			"user_id":     submit.UserID,
			"nick":        submit.Nick,
			"problem_id":  submit.ProblemID,
			"contest_id":  submit.ContestID,
			"source_code": submit.SourceCode,
			"language":    helper.LanguageType(submit.Language),
			"status":      submit.Status,
			"time":        submit.Time,
			"memory":      submit.Memory,
			"submit_time": submit.SubmitTime,
		})
	}

	return submitsData
}

func (model *Submit) GetUserSubmits(userID uint) helper.ReturnType {
	var submits []Submit
	err := db.
		Model(&Submit{}).
		Select(submitFields).
		Order("id desc").
		Where("user_id = ?", userID).
		Find(&submits).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取提交记录失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取提交记录成功", Data: submits}
	}
}

func (model *Submit) AddSubmit(submit *Submit) helper.ReturnType {

	err := db.
		Omit("time").
		Create(submit).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "添加提交记录失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加提交记录成功", Data: ""}
}

func (model *Submit) UpdateStatusAfterSubmit(id int, data map[string]interface{}) helper.ReturnType {
	err := db.
		Model(&Submit{}).
		Where("id = ?", id).
		Select([]string{"status", "time", "memory", "msg"}).
		Updates(data).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: 0}
	}
}

func (model *Submit) GetAllSubmit(Offset int, Limit int, UserId uint) helper.ReturnType {
	var submits []Submit
	var count int

	err := db.
		Model(&Submit{}).
		Select(submitFields).
		Where("user_id = ?", UserId).
		Count(&count).
		Limit(Limit).
		Offset(Offset).
		Order("id desc").
		Find(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询提交记录失败", Data: err.Error()}
	}

	submitsData := model.GetReturnData(submits)

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询提交记录成功", Data: gin.H{
		"data":  submitsData,
		"count": count,
	}}

}

func (model *Submit) GetContestSubmits(ContestID uint) helper.ReturnType {
	var submits []Submit
	fields := []string{"user_id", "problem_id", "status", "submit_time"}
	err := db.Select(fields).Where("contest_id = ?", ContestID).Find(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询比赛提交记录失败", Data: nil}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询比赛提交记录成功", Data: submits}
}

// TODO
func (model *Submit) GetContestSubmitByUser(UserID uint, ContestID uint, PageNumber int) helper.ReturnType {
	var submits []Submit
	var count int

	err := db.
		Model(&Submit{}).
		Order("submit_time").
		Order("id desc").
		Where("contest_id = ? AND user_id = ?", ContestID, UserID).
		Count(&count).
		Offset((PageNumber - 1) * constants.PageSubmitLogLimit).
		Limit(constants.PageSubmitLogLimit).
		Find(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询提交记录失败", Data: err.Error()}
	}

	submitsData := model.GetReturnData(submits)

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询提交记录成功", Data: gin.H{
		"data":  submitsData,
		"count": count,
	}}

}

func (model *Submit) GetProblemSubmit(submit Submit) helper.ReturnType {
	data := Submit{}
	err := db.
		Order("id desc").
		Where("problem_id = ? and user_id = ?", submit.ProblemID, submit.UserID).
		Last(&data).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询提交记录失败", Data: err.Error()}
	}

	submitData := model.GetReturnData(append([]Submit{}, data))
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询提交记录成功", Data: submitData[0]}
}

func (model *Submit) GetContestSubmitsByTime(contestID uint, beginTime, endTime time.Time) helper.ReturnType {
	var submits []Submit

	err := db.Where("contest_id = ? AND submit_time BETWEEN ? AND ?", contestID, beginTime, endTime).Order("id").Find(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取比赛提交失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取比赛提交成功", Data: submits}
}

func (model *Submit) GetSubmitByID(id uint, UserID uint) helper.ReturnType {
	var submit Submit
	err := db.Where("id = ? and user_id = ?", id, UserID).
		Find(&submit).
		Error

	status := submit.Status

	if status != "CE" && status != "SE" && status != "UE" {
		submit.Msg = ""
	}

	submitData := map[string]interface{}{
		"id":          submit.ID,
		"user_id":     submit.UserID,
		"nick":        submit.Nick,
		"problem_id":  submit.ProblemID,
		"contest_id":  submit.ContestID,
		"source_code": submit.SourceCode,
		"language":    helper.LanguageType(submit.Language),
		"status":      submit.Status,
		"time":        submit.Time,
		"msg":         submit.Msg,
		"memory":      submit.Memory,
		"submit_time": submit.SubmitTime,
	}

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取提交记录失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取提交记录成功", Data: submitData}
}

func (model *Submit) GetSubmitInPrintRequest(id uint, UserID uint) helper.ReturnType {
	var submit Submit
	err := db.
		Select([]string{"id", "user_id"}).
		Where("id = ? and user_id = ?", id, UserID).
		Find(&submit).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取提交记录失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取提交记录成功", Data: submit}
}

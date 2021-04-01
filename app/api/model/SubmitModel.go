package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"time"
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

func (Submit) TableName() string {
	return "submit"
}

func (model *Submit) GetUserSubmits(userID uint) helper.ReturnType {
	var submits []Submit
	err := db.Model(&Submit{}).Where("user_id = ?", userID).
		Find(&submits).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取提交记录失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取提交记录成功", Data: submits}
	}
}

func (model *Submit) AddSubmit(submit Submit) helper.ReturnType {

	err := db.Omit("time").Create(&submit).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加提交记录失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加提交记录成功", Data: ""}
}

func (model *Submit) UpdateStatusAfterSubmit(id int, data map[string]interface{}) helper.ReturnType {
	err := db.Model(&Submit{}).
		Where("id = ?", id).
		Select([]string{"status", "time", "memory", "msg"}).
		Updates(data).
		Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "更新失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "更新成功", Data: 0}
	}
}

func (model *Submit) GetAllSubmit(Offset int, Limit int) helper.ReturnType {
	var submits []Submit

	err := db.Offset(Offset).Limit(Limit).Last(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询提交记录失败", Data: err.Error()}
	}

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

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询提交记录成功", Data: submitsData}

}

// TODO
func (model *Submit) GetContestSubmit(UserID uint, ContestID uint, PageNumber int) helper.ReturnType {
	var submits []Submit

	err := db.Order("submit_time").
		Where("contest_id = ? AND user_id = ?", ContestID, UserID).
		Offset((PageNumber - 1) * common.PageSubmitLogLimit).
		Limit(common.PageSubmitLogLimit).
		First(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询提交记录失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询提交记录成功", Data: submits}

}

func (model *Submit) GetProblemSubmit(submit Submit) helper.ReturnType {
	data := Submit{}
	err := db.Where("problem_id = ? and user_id = ?", submit.ProblemID, submit.UserID).Last(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询提交记录失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询提交记录成功", Data: data}
}

func (model *Submit) GetContestSubmitsByTime(contestID uint, beginTime, endTime time.Time) helper.ReturnType {
	var submits []Submit

	err := db.Where("contest_id = ? AND submit_time BETWEEN ? AND ?", contestID, beginTime, endTime).Order("id").Find(&submits).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取比赛提交失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取比赛提交成功", Data: submits}
}

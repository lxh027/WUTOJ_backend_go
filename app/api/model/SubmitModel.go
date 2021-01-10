package model

import (
	"OnlineJudge/app/helper"
	"time"
)

type Submit struct {
	ID 			uint 		`json:"id" form:"id"`
	UserID 		uint 		`json:"user_id" form:"user_id"`
	Nick		string 		`json:"nick" form:"nick"`
	ProblemID	uint 		`json:"problem_id" form:"problem_id"`
	ContestID	uint 		`json:"contest_id" form:"contest_id"`
	SourceCode	string 		`json:"source_code" form:"source_code"`
	Language	int 		`json:"language" form:"language"`
	Status  	string 		`json:"status" form:"status"`
	Time  		int64 		`json:"time" form:"time"`
	Memory		uint 		`json:"memory" form:"memory"`
	SubmitTime	time.Time 	`json:"submit_time" form:"submit_time"`
}

func (Submit) TableName() string {
	return "submit"
}

func (model *Submit) GetUserSubmits(userID uint) helper.ReturnType {
	submits := make([]Submit, 0)
	err := db.Select([]string{"status", "count(*) as cnt"}).
		Where("user_id = ?", userID).
		Group("status").
		Find(&submits).Error
	if err != nil {
		return helper.ReturnType{Status: helper.CODE_ERROE, Msg: "获取提交记录失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: helper.CODE_SUCCESS, Msg: "获取提交记录成功", Data: submits}
	}
}
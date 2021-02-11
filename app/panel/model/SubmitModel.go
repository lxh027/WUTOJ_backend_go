package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"time"
)

type Submit struct {
	ID 		int 	`json:"id" form:"id"`
	UserID 	int		`json:"user_id" form:"user_id"`
	Nick 	string 	`json:"nick" form:"nick"`
	ProblemID 	int `json:"problem_id" form:"problem_id"`
	ContestID 	int `json:"contest_id" form:"contest_id"`
	SourceCode 	string `json:"source_code" form:"source_code"`
	Language 	int 	`json:"language" form:"language"`
	Status 		string 	`json:"status" form:"status"`
	Msg 		string 	`json:"msg" form:"msg"`
	Time 		int 	`json:"time" form:"time"`
	Memory 		int 	`json:"memory" form:"memory"`
	SubmitTime 	time.Time	`json:"submit_time" form:"submit_time"`
}


func (model *Submit) GetAllSubmit(offset int, limit int, whereData map[string]string, minSubmitTime, maxSubmitTime time.Time) helper.ReturnType {
	var submits []Submit
	/*where := "user_id like ? AND problem_id like ? AND contest_id like ? AND language like ? " +
		"AND status like ? AND submit_time >= ? AND submit_time <= ?"*/
	where := ""
	var count int
	var args []interface{}

	for field, data := range whereData {
		if data != "" {
			where += field + " = ? AND "
			args = append(args, data)
		}
	}

	args = append(args, minSubmitTime, maxSubmitTime)
	where += "submit_time >= ? AND submit_time <= ?"

	db.Model(&Submit{}).
		Where(where, args...).
		Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, args...).
		Order("id desc").
		Find(&submits).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"submits": submits,
				"count": count,
			},
		}
	}
}


func (model *Submit) GetSubmitGroup(whereData map[string]string, minSubmitTime, maxSubmitTime time.Time) helper.ReturnType {
	var submits []Submit
	where := ""
	var args []interface{}
	for field, data := range whereData {
		if data != "" {
			where += field + " = ? AND "
			args = append(args, data)
		}
	}
	args = append(args, minSubmitTime, maxSubmitTime)
	where += "submit_time >= ? AND submit_time <= ?"

	err := db.
		Where(where, args...).
		Find(&submits).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: submits}
	}
}

func (model *Submit) FindSubmitByID(id int) helper.ReturnType {
	var submit Submit

	err := db.Where("id = ?", id).First(&submit).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: submit}
	}
}

func (model *Submit) UpdateStatusAfterSubmit(id int, data map[string]interface{}) helper.ReturnType {
	err := db.Model(&Submit{}).
		Select([]string{"status", "time", "memory", "msg"}).
		Updates(data).
		Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "更新失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "更新成功", Data: 0}
	}
}

package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
)

type Sample struct {
	SampleID  int    `json:"sample_id" form:"sample_id"`
	ProblemID int    `json:"problem_id" form:"problem_id"`
	Input     string `json:"input" form:"input"`
	Output    string `json:"output" form:"output"`
}

func (model *Sample) FindSamplesByProblemID(id int) helper.ReturnType {
	var samples []Sample

	err := db.Where("problem_id = ?", id).Find(&samples).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: samples}
	}
}

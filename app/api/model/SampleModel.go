package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Sample struct {
	SampleID  int    `json:"sample_id" form:"sample_id"`
	ProblemID int    `json:"problem_id" form:"problem_id"`
	Input     string `json:"input" form:"input"`
	Output    string `json:"output" form:"output"`
}

func (model *Sample) FindSamplesByProblemID(id int) helper.ReturnType {
	var samples []Sample

	err := db.
		Select([]string{"sample_id", "problem_id", "input", "output"}).
		Where("problem_id = ?", id).
		Find(&samples).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: samples}
	}
}

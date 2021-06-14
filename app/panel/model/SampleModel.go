package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Sample struct {
	SampleID 	int 	`json:"sample_id" form:"sample_id"`
	ProblemID 	int 	`json:"problem_id" form:"problem_id"`
	Input 		string 	`json:"input" form:"input"`
	Output 		string 	`json:"output" form:"output"`
}

func (model *Sample) AddSample(newSample Sample) helper.ReturnType {//jun
	err := db.Create(&newSample).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *Sample) DeleteSample(sampleID int) helper.ReturnType  {
	err := db.Where("sample_id = ?", sampleID).Delete(Sample{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Sample) UpdateSample(sampleID int, updateSample Sample) helper.ReturnType  {
	err := db.Model(&Sample{}).Where("sample_id = ?", sampleID).Update(updateSample).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Sample) FindSamplesByProblemID(id int) helper.ReturnType {
	var samples []Sample

	err := db.Where("problem_id = ?", id).Find(&samples).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: samples}
	}
}
package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
)

type Problem struct {
	ProblemID    uint    `json:"problem_id" form:"problem_id"`
	Title        string  `json:"title" form:"title" `
	Background   string  `json:"background" form:"background"`
	Describe     string  `json:"describe" form:"describe"`
	InputFormat  string  `json:"input_format" form:"input_format"`
	OutputFormat string  `json:"output_format" form:"output_format"`
	Hint         string  `json:"hint" form:"hint"`
	Public       uint    `json:"public" form:"public"`
	Source       string  `json:"source" form:"source"`
	Time         float64 `json:"time" form:"time"`
	Memory       int     `json:"memory" form:"memory"`
	Type         string  `json:"type" form:"type"`
	Tag          string  `json:"tag" form:"tag"`
	Path         string  `json:"path" form:"path"`
	Status       int     `json:"status" form:"status"`
}

func (model *Problem) GetAllProblems() helper.ReturnType {
	var Problems []Problem
	var count int

	db.Model(&Problem{}).Count(&count)

	err := db.Find(&Problems).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"data":  Problems,
				"count": count,
			},
		}
	}

}

func (model *Problem) GetProblemByID(id int) helper.ReturnType {
	var problem Problem

	err := db.Where("problem_id = ?", id).First(&problem).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: problem}
	}
}

func (model *Problem) GetProblemByTitle(title string) helper.ReturnType {
	var problem Problem

	err := db.Where("title = ?", title).First(&problem).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: problem}
	}
}

func (model *Problem) SearchProblem(param string) helper.ReturnType {
	problem := Problem{}
	err := db.Where("problem_id = ?", param).Find(&problem).Error

	if err == nil {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: problem}
	} else {
		if err = db.Where("title = ?", param).Find(&problem).Error; err == nil {
			return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: problem}
		} else {
			if err = db.Where("tag = ?", param).Find(&problem).Error; err == nil {
				return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: problem}
			}
		}
	}

	return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: ""}

}
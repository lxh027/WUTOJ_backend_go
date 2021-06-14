package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Problem struct {
	ProblemID 	int 	`json:"problem_id" form:"problem_id"`
	Title 		string 	`json:"title" form:"title"`
	Background	string 	`json:"background" form:"background"`
	Describe 	string 	`json:"describe" form:"describe"`
	InputFormat string 	`json:"input_format" form:"input_format"`
	OutputFormat string `json:"output_format" form:"output_format"`
	Hint 		string 	`json:"hint" form:"hint"`
	Public 		int 	`json:"public" form:"public"`
	Source 		string 	`json:"source" form:"source"`
	Time 		float32	`json:"time" form:"time"`
	Memory 		int 	`json:"memory" form:"memory"`
	Type  		string 	`json:"type" form:"type"`
	Tag 		string 	`json:"tag" form:"tag"`
	Path 		string 	`json:"path" form:"path"`
	Status 		int 	`json:"status" form:"status"`
}


func (model *Problem) GetAllProblem(offset int, limit int, title string) helper.ReturnType {
	var problems []Problem
	where := "title like ?"
	var count int

	db.Model(&Problem{}).Where(where, "%"+title+"%").Count(&count)


	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+title+"%").
		Find(&problems).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"problems": problems,
				"count": count,
			},
		}
	}
}

func (model *Problem) FindProblemByID(id int) helper.ReturnType {
	var problem Problem

	err := db.Where("problem_id = ?", id).First(&problem).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: problem}
	}
}

func (model *Problem) AddProblem(newProblem Problem) helper.ReturnType {//jun
	err := db.Create(&newProblem).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		var idProblem Problem
		db.First(&idProblem, newProblem)
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: idProblem.ProblemID}
	}
}

func (model *Problem) DeleteProblem(problemID int) helper.ReturnType  {
	err := db.Where("problem_id = ?", problemID).Delete(Problem{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Problem) UpdateProblem(problemID int, updateProblem Problem) helper.ReturnType  {
	err := db.Model(&Problem{}).Where("problem_id = ?", problemID).Update(updateProblem).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Problem) ChangeProblemStatus(problemID int, status int) helper.ReturnType  {
	err := db.Model(&Problem{}).Where("problem_id = ?", problemID).Update("status", status).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Problem) ChangeProblemPublicStatus(problemID int, isPublic int) helper.ReturnType  {
	err := db.Model(&Problem{}).Where("problem_id = ?", problemID).Update("public", isPublic).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Problem) SaveProblemPath(problemID int, path string) helper.ReturnType {
	err := db.Model(&Problem{}).Where("problem_id = ?", problemID).Update("path", path).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}

}
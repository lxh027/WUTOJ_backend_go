package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Problem struct {
	ProblemID    uint    `json:"problem_id,omitempty" form:"problem_id"`
	Title        string  `json:"title" form:"title" `
	Background   string  `json:"background" form:"background"`
	Describe     string  `json:"describe" form:"describe"`
	InputFormat  string  `json:"input_format" form:"input_format"`
	OutputFormat string  `json:"output_format" form:"output_format"`
	Hint         string  `json:"hint" form:"hint"`
	Public       uint    `json:"public,omitempty" form:"public"`
	Source       string  `json:"source" form:"source"`
	Time         float64 `json:"time" form:"time"`
	Memory       int     `json:"memory" form:"memory"`
	Type         string  `json:"type" form:"type"`
	Tag          string  `json:"tag" form:"tag"`
	Path         string  `json:"path,omitempty" form:"path"`
	Status       int     `json:"status,omitempty" form:"status"`
}

// []string{"problem_id", "time", "title", "background", "describe", "input_format", "output_format", "hint", "public", "source", "memory", "type", "tag", "path", "status"}
func (model *Problem) GetAllProblems(offset int, limit int) helper.ReturnType {
	var Problems []Problem
	var count int

	err := db.
		Select([]string{"public"}).
		Model(&Problem{}).
		Where("public = ?", constants.ProblemPublic).
		Count(&count).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}

	err = db.
		Select([]string{"problem_id", "time", "title", "background", "`describe`", "input_format", "output_format", "hint", "public", "source", "memory", "type", "tag", "status"}).
		Offset(offset).
		Limit(limit).
		Where("public = ?", constants.ProblemPublic).
		Order("problem_id").
		Find(&Problems).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		var problemData []map[string]interface{}
		sampleModel := Sample{}
		problemSubmitLog := ProblemSubmitLog{}
		for _, problem := range Problems {
			res := sampleModel.FindSamplesByProblemID(int(problem.ProblemID))
			problemData = append(problemData, map[string]interface{}{
				"problem":            problem,
				"problem_sample":     res.Data,
				"problem_submit_log": problemSubmitLog.GetProblemSubmitLog(problem.ProblemID).Data,
			})
		}

		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"data":  problemData,
				"count": count,
			},
		}
	}

}

func (model *Problem) GetProblemFieldsByIDList(ids []uint, fields []string) helper.ReturnType {
	var problems []Problem

	err := db.
		Select(fields).
		Where("problem_id in (?)", ids).
		Find(&problems).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: problems}
}

// Use
func (model *Problem) GetProblemByID(id int) helper.ReturnType {
	var problem Problem
	sampleModel := Sample{}
	problemSubmitLog := ProblemSubmitLog{}

	err := db.
		Select([]string{"problem_id", "time", "title", "background", "`describe`", "input_format", "output_format", "hint", "source", "memory", "type", "tag", "public"}).
		Where("problem_id = ?", id).
		First(&problem).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		problemData := map[string]interface{}{
			"problem":            problem,
			"problem_sample":     sampleModel.FindSamplesByProblemID(int(problem.ProblemID)).Data,
			"problem_submit_log": problemSubmitLog.GetProblemSubmitLog(problem.ProblemID).Data,
		}
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: problemData}
	}
}

// Unused
func (model *Problem) GetProblemByTitle(title string) helper.ReturnType {
	var problem Problem

	err := db.
		Where("title = ? AND public = ?", title, constants.ProblemPublic).
		First(&problem).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: problem}
	}
}

func (model *Problem) SearchProblem(param string) helper.ReturnType {
	problem := Problem{}
	err := db.
		Where("problem_id = ?", param).
		Find(&problem).
		Error

	sampleModel := Sample{}
	problemSubmitLog := ProblemSubmitLog{}
	if err == nil {

		problemData := map[string]interface{}{
			"problem":            problem,
			"problem_sample":     sampleModel.FindSamplesByProblemID(int(problem.ProblemID)).Data,
			"problem_submit_log": problemSubmitLog.GetProblemSubmitLog(problem.ProblemID).Data,
		}

		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: problemData}
	}

	return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
}

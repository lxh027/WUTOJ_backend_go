package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type ContestProblem struct {
	ContestID 	int 	`json:"contest_id" form:"contest_id"`
	ProblemID 	int 	`json:"problem_id" form:"problem_id"`
}

func (model *ContestProblem) GetContestProblems(contestID int) helper.ReturnType {
	var contestProblems []ContestProblem

	err := db.Where("contest_id = ?", contestID).Find(&contestProblems).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "get error", Data: err.Error()}
	}

	var problems []int
	for _, item := range contestProblems {
		problems = append(problems, item.ProblemID)
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "success", Data: problems}
}

func (model *ContestProblem) AddContestProblem(newContestProblem ContestProblem) helper.ReturnType {
	err := db.Create(&newContestProblem).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "add problem contest error", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "success", Data: nil}
}

func (model *ContestProblem) DeleteContestProblem(contestID int) helper.ReturnType  {
	err := db.Where("contest_id = ?", contestID).Delete(&ContestProblem{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "delete error", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "success", Data: nil}
}
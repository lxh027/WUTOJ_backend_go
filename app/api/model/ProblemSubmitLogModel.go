package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type ProblemSubmitLog struct {
	ProblemID int `json:"problem_id" form:"problem_id"`
	AC        int `json:"ac" form:"ac"`
	WA        int `json:"wa" form:"wa"`
	TLE       int `json:"tle" form:"tle"`
	MLE       int `json:"mle" form:"mle"`
	RE        int `json:"re" form:"re"`
	SE        int `json:"se" form:"se"`
	CE        int `json:"ce" form:"ce"`
}

func (*ProblemSubmitLog) GetProblemSubmitLog(ProblemID uint) helper.ReturnType {
	problemSubmitLog := ProblemSubmitLog{}

	err := db.
		Select([]string{"problem_id", "ac", "wa", "tle", "mle", "re", "se", "ce"}).
		Where("problem_id = ?", ProblemID).
		Find(&problemSubmitLog).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询题目提交记录失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询题目提交记录成功", Data: problemSubmitLog}
}

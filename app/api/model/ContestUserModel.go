package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
)

type ContestUser struct {
	ContestID int `json:"contest_id" form:"contest_id"`
	UserID    int `json:"user_id" form:"user_id"`
	ID        int `json:"id" form:"id"`
	Status    int `json:"status" form:"status"`
}

func (model *ContestUser) AddContestUser(data ContestUser) helper.ReturnType {
	err := db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "参加比赛失败", Data: ""}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "已参加比赛", Data: ""}
	}
}

func (model *ContestUser) GetUserContest(UserID int) helper.ReturnType {
	// to do
	// ask ljw for detail
}

package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type ContestUser struct {
	ContestID int `json:"contest_id" form:"contest_id" uri:"contest_id"`
	UserID    int `json:"user_id" form:"user_id"`
	ID        int `json:"id" form:"id"`
	Status    int `json:"status" form:"status"`
}

func (ContestUser) TableName() string {
	return "contest_users"
}

func (model *ContestUser) AddContestUser(data ContestUser) helper.ReturnType {
	contestUser := ContestUser{}
	err := db.
		Select([]string{"contest_id", "user_id"}).
		Where("contest_id = ? AND user_id = ?", data.ContestID, data.UserID).
		Find(&contestUser).
		Error
	if err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "已经参加比赛，请勿重复参赛", Data: ""}
	}

	contest := Contest{}
	err = db.Select([]string{"rule"}).
		Where("contest_id = ?", data.ContestID).
		Find(&contest).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "比赛不存在", Data: ""}
	}

	if contest.Rule != constants.CommonContest {
		return helper.ReturnType{Status: constants.CodeError, Msg: "比赛不可自己报名", Data: ""}
	}

	err = db.Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "参加比赛失败", Data: ""}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "参加比赛成功", Data: ""}
	}
}

func (model *ContestUser) CheckUserContest(UserID int, ContestID int) helper.ReturnType {

	contestUser := ContestUser{}

	err := db.
		Select([]string{"contest_id", "user_id"}).
		Where("user_id = ? AND contest_id = ?", UserID, ContestID).
		Find(&contestUser).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "还未参加比赛，请参加比赛", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: contestUser}

}

func (model *ContestUser) GetUserContest(UserID int) helper.ReturnType {

	var contestUser []ContestUser

	err := db.
		Select([]string{"user_id", "contest_id"}).
		Where("user_id = ?", UserID).
		Find(&contestUser).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询用户比赛成功", Data: contestUser}

}

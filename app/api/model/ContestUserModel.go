package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
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

	if err := db.Where("contest_id = ?", data.ContestID).Where("user_id = ?", data.UserID).Find(&contestUser).Error; err == nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "已经参加比赛，请勿重复参赛", Data: ""}
	}

	err := db.Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "参加比赛失败", Data: ""}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "参加比赛成功", Data: ""}
	}
}

func (model *ContestUser) GetUserContest(UserID int) helper.ReturnType {

	contestUser := ContestUser{}

	err := db.Where("user_id = ?", UserID).Find(&contestUser).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: contestUser}

}

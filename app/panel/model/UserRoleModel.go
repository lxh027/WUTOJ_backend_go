package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type UserRole struct {
	UserID int `json:"user_id" form:"user_id"`
	Rid    int `json:"rid" form:"rid"`
}

func (model *UserRole) AddUserRole(newUserRole UserRole) helper.ReturnType {
	err := db.Create(&newUserRole).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *UserRole) DeleteUserRole(newUserRole UserRole) helper.ReturnType {
	err := db.Where("user_id = ? AND rid = ?", newUserRole.UserID, newUserRole.Rid).Delete(UserRole{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

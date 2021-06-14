package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type RoleAuth struct {
	Rid 	int 	`json:"rid" form:"rid"`
	Aid 	int 	`json:"aid" form:"aid"`
}

func (model *RoleAuth) AddRoleAuth(newRoleAuth RoleAuth) helper.ReturnType {
	err := db.Create(&newRoleAuth).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}
func (model *RoleAuth) DeleteRoleAuth(newRoleAuth RoleAuth) helper.ReturnType {
	err := db.Where("rid = ? AND aid = ?", newRoleAuth.Rid, newRoleAuth.Aid).Delete(RoleAuth{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}
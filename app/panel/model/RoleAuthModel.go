package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
)

type RoleAuth struct {
	Rid 	int 	`json:"rid" form:"rid"`
	Aid 	int 	`json:"aid" form:"aid"`
}

func (model *RoleAuth) AddRoleAuth(newRoleAuth RoleAuth) helper.ReturnType {
	err := db.Create(&newRoleAuth).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}
}
func (model *RoleAuth) DeleteRoleAuth(newRoleAuth RoleAuth) helper.ReturnType {
	err := db.Where("rid = ? AND aid = ?", newRoleAuth.Rid, newRoleAuth.Aid).Delete(RoleAuth{}).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除成功", Data: true}
	}
}
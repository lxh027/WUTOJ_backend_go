package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

//OJWebUserConfig 用于oj网站和用户的配置表
type OJWebUserConfig struct {
	ID             int    `json:"id" form:"id"`
	OJID           int    `json:"oj_id" form:"oj_id"`
	UserID         int    `json:"user_id" form:"user_id"`
	OJUserName     string `json:"oj_user_name" form:"oj_user_name"`
	OJUserPassword string `json:"oj_user_password" form:"oj_user_password"`
	ACProblems     string `json:"ac_problems" form:"ac_problems"`
	Status         int    `json:"status" form:"status"`
}

//TableName 设定表名
func (OJWebUserConfig) TableName() string {
	return "oj_web_user_config"
}

//AddOJWebUserConfig 添加用户OJ配置
func (model *OJWebUserConfig) AddOJWebUserConfig(newOJWebUserConfig OJWebUserConfig) helper.ReturnType {
	ojWebUserConfig := OJWebUserConfig{}
	if err := db.Where("nick = ? AND mail = ?", newOJWebUserConfig.UserID, newOJWebUserConfig.OJID).First(&ojWebUserConfig).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户配置已存在", Data: false}
	}

	err := db.Create(&newOJWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
}

//DeleteOJWebUserConfig 删除用户oj配置
func (model *OJWebUserConfig) DeleteOJWebUserConfig(userID int, ojID int) helper.ReturnType {
	err := db.Where("user_id = ? AND oj_id = ?", userID, ojID).Delete(OJWebUserConfig{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
}

//UpdateOJWebUserConfig 修改用户oj配置
func (model *OJWebUserConfig) UpdateOJWebUserConfig(userID int, ojID int, updateOJWebUserConfig OJWebUserConfig) helper.ReturnType {
	err := db.Model(&User{}).Where("user_id = ? AND oj_id = ?", userID, ojID).Update(updateOJWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
}

//TODO:delete this later
// func (model *OJWebUserConfig) GetOJWebUserConfigByUserID(userID int) helper.ReturnType {
// 	var getUser User

// 	err := db.Select([]string{"user_id", "id", "oj_id", "oj_user_name", "oj_user_password", "ac_problems", "status"}).Where("user_id = ?", userID).First(&getUser).Error
// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	}
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: getUser}
// }

//GetUserOJwebUserConfig 获取用户oj网站配置信息
func (model *OJWebUserConfig) GetUserOJwebUserConfig(UserID int) helper.ReturnType {

	var ojWebUserConfig []OJWebUserConfig

	err := db.Where("user_id = ?", UserID).Find(&ojWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询用户比赛成功", Data: ojWebUserConfig}

}

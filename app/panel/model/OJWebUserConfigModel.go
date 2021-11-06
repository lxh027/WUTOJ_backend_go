package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

//OJWebUserConfig 用于oj网站和用户的配置表
type OJWebUserConfig struct {
	ID         int    `json:"id" form:"id"`
	OJName     string `json:"oj_name" form:"oj_name"`
	UserID     int    `json:"user_id" form:"user_id"`
	OJUserName string `json:"oj_user_name" form:"oj_user_name"`
	Status     int    `json:"status" form:"status"`
}

//TableName 设定表名
func (OJWebUserConfig) TableName() string {
	return "oj_web_user_config"
}

//AddOJWebUserConfig 添加用户OJ配置
func (model *OJWebUserConfig) AddOJWebUserConfig(newOJWebUserConfig OJWebUserConfig) helper.ReturnType {
	ojWebUserConfig := OJWebUserConfig{}
	if err := db.Where("user_id = ? AND oj_name = ?", newOJWebUserConfig.UserID, newOJWebUserConfig.OJName).First(&ojWebUserConfig).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户配置已存在", Data: false}
	}

	err := db.Create(&newOJWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
}

//DeleteOJWebUserConfig 删除用户oj配置
func (model *OJWebUserConfig) DeleteOJWebUserConfig(id int) helper.ReturnType {
	err := db.Where("id = ?", id).Delete(OJWebUserConfig{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
}

//UpdateOJWebUserConfig 修改用户oj配置
func (model *OJWebUserConfig) UpdateOJWebUserConfig(id int, updateOJWebUserConfig OJWebUserConfig) helper.ReturnType {
	err := db.Model(&OJWebUserConfig{}).Where("id = ?", id).Update(updateOJWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
}

//GetUserOJwebUserConfig 获取用户oj网站配置信息
func (model *OJWebUserConfig) GetUserOJwebUserConfig(userID int) helper.ReturnType {

	var ojWebUserConfig []OJWebUserConfig

	err := db.Where("user_id = ?", userID).Find(&ojWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询用户比赛成功", Data: ojWebUserConfig}

}

//GetOJWebUserConfigByID 由ID获取OJ用户配置
func (model *OJWebUserConfig) GetOJWebUserConfigByID(id int) helper.ReturnType {

	var ojWebUserConfig OJWebUserConfig

	err := db.Where("id = ?", id).Find(&ojWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询用户比赛成功", Data: ojWebUserConfig}

}

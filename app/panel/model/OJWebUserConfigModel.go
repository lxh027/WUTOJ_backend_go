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

type OJWebUserConfigWithNick struct {
	ID         int    `json:"id" form:"id"`
	OJName     string `json:"oj_name" form:"oj_name"`
	UserID     int    `json:"user_id" form:"user_id"`
	OJUserName string `json:"oj_user_name" form:"oj_user_name"`
	Status     int    `json:"status" form:"status"`
	Nick       string `json:"nick" form:"nick"`
	Realname   string `json:"realname" form:"realname"`
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
func (model *OJWebUserConfig) GetUserOJwebUserConfig(userID int, ojName string) helper.ReturnType {

	var ojWebUserConfig []OJWebUserConfig

	err := db.Where("user_id = ? AND oj_name = ?", userID, ojName).Find(&ojWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询用户OJ配置成功", Data: ojWebUserConfig}

}

//GetOJWebUserConfigByID 由ID获取OJ用户配置
func (model *OJWebUserConfig) GetOJWebUserConfigByID(id int) helper.ReturnType {

	var ojWebUserConfig OJWebUserConfig

	err := db.Where("id = ?", id).Find(&ojWebUserConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询用户OJ配置成功", Data: ojWebUserConfig}

}

//GetAllOJWebUserConfig 获取所有用户OJ配置
func (model *OJWebUserConfig) GetAllOJWebUserConfig(offset int, limit int, ojName string) helper.ReturnType {
	var ojWebUserConfigs []OJWebUserConfigWithNick
	where := "oj_name like ?"
	var count int

	db.Model(&OJWebUserConfig{}).
		Where(where, "%"+ojName+"%").
		Count(&count)

	err := db.Table("oj_web_user_config").
		Select("oj_web_user_config.id as id,oj_web_user_config.oj_name as oj_name,oj_web_user_config.user_id as user_id,oj_web_user_config.oj_user_name as oj_user_name,oj_web_user_config.status as status,users.nick as nick,users.realname as realname").
		Joins("LEFT JOIN users ON users.user_id = oj_web_user_config.user_id").
		Offset(offset).
		Limit(limit).
		Where(where, "%"+ojName+"%").
		Scan(&ojWebUserConfigs).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"oj_web_user_configs": ojWebUserConfigs,
				"count":               count,
			},
		}
	}
}

//ChangeOJConfigStatus 变更配置状态
func (model *OJWebUserConfig) ChangeOJConfigStatus(id int, status int) helper.ReturnType {
	err := db.Model(&OJWebUserConfig{}).Where("id = ?", id).Update("status", status).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

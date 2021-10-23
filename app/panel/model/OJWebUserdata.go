package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

//OJWebUserData 储存从oj爬取的数据的数据表
type OJWebUserData struct {
	ID         int       `json:"id" form:"id"`
	OJID       int       `json:"oj_id" form:"oj_id"`
	UserID     int       `json:"user_id" form:"user_id"`
	SubmitTime time.Time `json:"submit_time" form:"submit_time"`
	Status     int       `json:"status" form:"status"`
}

//TableName 设定表名
func (OJWebUserData) TableName() string {
	return "oj_webs"
}

//AddOJWebUserData 添加用户OJ数据
func (model *OJWebUserConfig) AddOJWebUserData(newOJWebUserData OJWebUserData) helper.ReturnType {
	ojWebUserData := OJWebUserData{}
	if err := db.Where("user_id = ? AND submit_time = ?", newOJWebUserData.UserID, newOJWebUserData.SubmitTime).First(&ojWebUserData).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户配置已存在", Data: false}
	}

	err := db.Create(&newOJWebUserData).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
}

//DeleteOJWebUserData 删除某用户在特定OJ的所有数据
func (model *User) DeleteOJWebUserData(userID int, ojID int) helper.ReturnType {
	err := db.Where("user_id = ? AND oj_id = ?", userID, ojID).Delete(OJWebUserData{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
}

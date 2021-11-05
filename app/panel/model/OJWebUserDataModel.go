package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

//OJWebUserData 储存从oj爬取的数据的数据表
type OJWebUserData struct {
	ID         int       `json:"id" form:"id"`
	OJName     string    `json:"oj_name" form:"oj_name"`
	UserID     int       `json:"user_id" form:"user_id"`
	ProblemID  string    `json:"problem_id" form:"problem_id"`
	SubmitTime time.Time `json:"submit_time" form:"submit_time"`
	Status     string    `json:"status" form:"status"`
}

//TableName 设定表名
func (OJWebUserData) TableName() string {
	return "oj_web_data"
}

//AddOJWebUserData 添加用户OJ数据
func (model *OJWebUserData) AddOJWebUserData(newOJWebUserData OJWebUserData) helper.ReturnType {
	ojWebUserData := OJWebUserData{}
	if err := db.Where("user_id = ? AND submit_time = ?", newOJWebUserData.UserID, newOJWebUserData.SubmitTime).First(&ojWebUserData).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeSuccess /*这里视为正常添加*/, Msg: "用户提交记录已存", Data: false}
	}

	err := db.Create(&newOJWebUserData).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
}

//DeleteOJWebUserData 删除某用户在特定OJ的所有数据
func (model *OJWebUserData) DeleteOJWebUserData(userID int, ojID string) helper.ReturnType {
	err := db.Where("user_id = ? AND oj_name = ?", userID, ojID).Delete(OJWebUserData{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
}

//AddOJWebUserDatas 添加多处用户OJ数据
func (model *OJWebUserData) AddOJWebUserDatas(newOJWebUserDatas []OJWebUserData) helper.ReturnType {
	ojWebUserData := OJWebUserData{}
	tx := db.Begin()

	for _, newOJWebUserData := range newOJWebUserDatas {
		if err := tx.Where("user_id = ? AND submit_time = ?", newOJWebUserData.UserID, newOJWebUserData.SubmitTime).First(&ojWebUserData).Error; err == nil {
			continue
		}
		err := tx.Create(&newOJWebUserData).Error
		if err != nil {
			tx.Rollback()
			return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
		}
	}
	tx.Commit()
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}

}

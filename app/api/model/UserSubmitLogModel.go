package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
)

type UserSubmitLog struct {
	UserID int `gorm:"user_id"`
	AC     int `gorm:"ac"`
	WA     int `gorm:"wa"`
	TLE    int `gorm:"tle"`
	MLE    int `gorm:"mle"`
	RE     int `gorm:"re"`
	SE     int `gorm:"se"`
	CE     int `gorm:"ce"`
}

func (UserSubmitLog) TableName() string {
	return "user_submit_log"
}

func (model *UserSubmitLog) GetUserSubmitLog(UserID uint) helper.ReturnType {
	var UserSubmitLog UserSubmitLog
	err := db.Model(&UserSubmitLog).Where("user_id = ?", int(UserID)).First(&UserSubmitLog).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询提交数据失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询提交数据成功", Data: UserSubmitLog}
}

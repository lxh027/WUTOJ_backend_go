package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

type PrintLog struct {
	ID       int       `json:"id,omitempty" gorm:"id"`
	SubmitID int       `json:"submit_id,omitempty" gorm:"submit_id"`
	Status   int       `json:"status,omitempty" gorm:"status"`
	PrintAt  time.Time `json:"print_at,omitempty" gorm:"print_at"`
}

func (PrintLog) TableName() string {
	return "print_log"
}

func (model *PrintLog) AddPrintLog(log PrintLog) helper.ReturnType {
	err := db.Model(&PrintLog{}).Create(&log).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "添加失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加成功", Data: 0}
}

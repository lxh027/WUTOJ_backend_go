package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"time"
)

type PrintLog struct {
	ID       int       `gorm:"id"`
	SubmitID int       `gorm:"submit_id"`
	Status   int       `gorm:"status"`
	PrintAt  time.Time `gorm:"print_at"`
}

func (PrintLog) TableName() string {
	return "print_log"
}

func (model *PrintLog) AddPrintLog(log PrintLog) helper.ReturnType {
	err := db.Model(&PrintLog{}).Create(&log).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加成功", Data: 0}
}

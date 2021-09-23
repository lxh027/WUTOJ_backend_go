package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

type PrintLog struct {
	ID        int       `json:"id,omitempty" gorm:"id"`
	Status    int       `json:"status,omitempty" gorm:"status"`
	PrintAt   time.Time `json:"print_at,omitempty" gorm:"print_at"`
	RequestAt time.Time `gorm:"request_at"`
	UserNick  string    `json:"user_nick" gorm:"user_nick"`
	Code      string    `json:"code" gorm:"code"`
}

func (PrintLog) TableName() string {
	return "print_log"
}

func (model *PrintLog) AddPrintLog(log PrintLog) helper.ReturnType {
	err := db.
		Model(&PrintLog{}).
		Omit("print_at").
		Create(&log).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "添加失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加成功", Data: 0}
}

package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"github.com/gin-gonic/gin"
	"time"
)

type PrintLog struct {
	ID       int       `json:"print_id" gorm:"print_id"`
	SubmitID int       `json:"submit_id" gorm:"submit_id"`
	Status   int       `json:"status" gorm:"status"`
	PrintAt  time.Time `json:"print_at" gorm:"print_at"`
}

func (PrintLog) TableName() string {
	return "print_log"
}

func (model *PrintLog) UpdateStatusAfterPrint(PrintLogID int, data PrintLog) helper.ReturnType {
	err := db.Model(&PrintLog{}).
		Where("id = ?", PrintLogID).
		Select([]string{"id", "status", "print_at"}).
		Updates(data).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新状态失败，数据库错误", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新状态成功", Data: 0}
}

func (model *PrintLog) GetAllPrintLog(Offset int, Limit int) helper.ReturnType {
	var printlogs []PrintLog
	var Count int
	err := db.Model(&PrintLog{}).
		Order("id desc").
		Count(&Count).
		Offset(Offset).
		Limit(Limit).
		Find(&printlogs).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: 0}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: gin.H{
			"data":  printlogs,
			"count": Count,
		},
		}
	}
}

func (model *PrintLog) GetPrintLogByID(id string) helper.ReturnType {
	printLog := PrintLog{}
	err := db.Model(&PrintLog{}).
		Where("id = ?", id).
		Find(&printLog).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: printLog}
}

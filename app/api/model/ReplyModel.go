package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"time"
)

type Reply struct {
	ID        int       `json:"id" form:"id"`
	DiscussID int       `json:"discuss_id" form:"discuss_id"`
	UserID    int       `json:"user_id" form:"user_id"`
	Content   string    `json:"content" form:"content"`
	Time      time.Time `json:"time" form:"time" gorm:"omitempty"`
}

func (model *Reply) AddReply(data Reply) helper.ReturnType {

	data.Time = time.Now()

	err := db.Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加成功", Data: true}
	}

}

package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"github.com/gin-gonic/gin"
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
		return helper.ReturnType{Status: constants.CodeError, Msg: "添加失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加成功", Data: true}
	}

}

func (model *Reply) GetReplyByDiscussID(DiscussID int, Offset int, Limit int) helper.ReturnType {
	var reply []Reply
	var count int

	err := db.
		Model(&Reply{}).
		Order("time desc").
		Where("discuss_id = ?", DiscussID).
		Count(&count).
		Offset(Offset).
		Limit(Limit).
		Find(&reply).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: gin.H{
			"data":  reply,
			"count": count,
		}}
	}

}

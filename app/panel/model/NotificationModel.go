package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

//Notification 在比赛期间的公告
type Notification struct {
	ID         int       `json:"id" form:"id"`
	Title      string    `json:"title" form:"title"`
	Content    string    `json:"content" form:"content"`
	SubmitTime time.Time `json:"submit_time" form:"submit_time"`
	ModifyTime time.Time `json:"modify_time" form:"modify_time"`
	ContestID  int       `json:"contest_id" form:"contest_id"`
	UserID     int       `json:"user_id" form:"user_id"`
	Status     int       `json:"status" form:"status"`
	EndTime    int       `json:"end_time" form:"end_time"`
}

func (model *Notification) GetAllNotification(contestID int) helper.ReturnType {

	var notifications []Notification
	var count int
	err := db.Model(&Notification{}).Where("contest_id = ?", contestID).Order("id desc").Find(&notifications).Count(&count).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"notifications": notifications,
				"count":         count,
			},
		}
	}
}

func (model *Notification) GetNotificationByID(id int) helper.ReturnType {
	var notification Notification

	err := db.Where("id = ?", id).First(&notification).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取成功", Data: notification}
	}

}

func (model *Notification) AddNotification(newNotice Notification) helper.ReturnType { //jun
	err := db.Create(&newNotice).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *Notification) DeleteNotification(noticeID int) helper.ReturnType {
	err := db.Where("id = ?", noticeID).Delete(Notification{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Notification) UpdateNotification(noticeID int, updateNotice Notification) helper.ReturnType {
	err := db.Model(&Notification{}).Where("id = ?", noticeID).Update(updateNotice).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Notification) ChangeNotificationStatus(noticeID int, status int) helper.ReturnType {
	err := db.Model(&Notification{}).Where("id = ?", noticeID).Update("status", status).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

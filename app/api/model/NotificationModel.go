package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

type Notification struct {
	ID         int       `json:"id,omitempty" form:"id"`
	Title      string    `json:"title,omitempty" form:"title"`
	Content    string    `json:"content,omitempty" form:"content"`
	SubmitTime time.Time `json:"submit_time,omitempty" form:"submit_time"`
	ModifyTime time.Time `json:"modify_time,omitempty" form:"modify_time"`
	ContestID  int       `json:"contest_id,omitempty" form:"contest_id"`
	UserID     int       `json:"user_id,omitempty" form:"user_id"`
	Status     int       `json:"status,omitempty" form:"status"`
	EndTime    int       `json:"end_time,omitempty" form:"end_time"`
}

func (model *Notification) GetNotification(ContestID int, LastNotification int) (helper.ReturnType, int) {
	var notifications []Notification

	err := db.
		Select([]string{"id", "content", "title", "submit_time", "modify_time"}).
		Where("contest_id = ?", ContestID).
		Where("status = ?", 1).
		Where("id > ?", LastNotification).
		Find(&notifications).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取失败", Data: err.Error()}, LastNotification
	} else {
		NotificationID := LastNotification
		for _, notify := range notifications {
			if notify.ID > NotificationID {
				NotificationID = notify.ID
			}
		}
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取成功", Data: notifications}, NotificationID
	}

}

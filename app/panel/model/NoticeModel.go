package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

//Notice 公告
type Notice struct {
	ID        int       `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	Link      string    `json:"link" form:"link"`
	Begintime time.Time `json:"begintime" form:"begintime" gorm:"column:begintime"`
	Endtime   time.Time `json:"endtime" form:"endtime" gorm:"column:endtime"`
}

func (model *Notice) GetAllNotice(offset int, limit int, title string, time time.Time) helper.ReturnType {
	var notices []Notice
	where := "title like ? AND begintime > ?"
	var count int

	db.Model(&Notice{}).Where(where, "%"+title+"%", time).Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+title+"%", time).
		Find(&notices).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"notices": notices,
				"count":   count,
			},
		}
	}
}

func (model *Notice) FindNoticeByID(id int) helper.ReturnType {
	var notice Notice

	err := db.Where("id = ?", id).First(&notice).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: notice}
	}
}

func (model *Notice) AddNotice(newNotice Notice) helper.ReturnType { //jun
	err := db.Create(&newNotice).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *Notice) DeleteNotice(noticeID int) helper.ReturnType {
	err := db.Where("id = ?", noticeID).Delete(Notice{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Notice) UpdateNotice(noticeID int, updateNotice Notice) helper.ReturnType {
	err := db.Model(&Notice{}).Where("id = ?", noticeID).Update(updateNotice).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

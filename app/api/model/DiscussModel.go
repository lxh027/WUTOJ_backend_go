package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"log"
	"time"
)

type Discuss struct {
	ID        int       `json:"id" form:"id"`
	ContestID int       `json:"contest_id" form:"contest_id"`
	ProblemID int       `json:"problem_id" form:"contest_id"`
	UserID    int       `json:"user_id" form:"user_id"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	Time      time.Time `json:"time" form:"time" gorm:"omitempty"`
	Status    int       `json:"status" form:"status"`
}

func (model *Discuss) GetAllDiscuss(Offset int, Limit int) helper.ReturnType {
	var discussions []Discuss

	err := db.Offset(Offset).Limit(Limit).Find(&discussions).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取失败，数据库错误", Data: ""}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: discussions}
	}
}

func (model *Discuss) GetDiscussionByID(id int, PageNumber int) helper.ReturnType {
	var discuss Discuss
	err := db.Where("id = ?", id).First(&discuss).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: ""}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: discuss}
	}

}

func (model *Discuss) AddDiscussion(newDiscussion Discuss) helper.ReturnType {
	//var discuss Discuss

	log.Print(newDiscussion)
	err := db.Omit("time").Create(&newDiscussion).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加成功", Data: true}
	}
}

func (model *Discuss) GetContestDiscussion(ContestID int, PageNumber int) helper.ReturnType {
	var discussions []Discuss

	err := db.Where("contest_id = ?", ContestID).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: discussions}
	}

}

func (model *Discuss) GetProblemDiscussion(ProblemID int, PageNumber int) helper.ReturnType {
	var discussions []Discuss

	err := db.Where("problem_id = ?", ProblemID).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: discussions}
	}
}

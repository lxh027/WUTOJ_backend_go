package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"

	"github.com/gin-gonic/gin"
)

type Discuss struct {
	ID        int       `json:"id,omitempty" form:"id"`
	ContestID int       `json:"contest_id,omitempty" form:"contest_id" uri:"contest_id"`
	ProblemID int       `json:"problem_id,omitempty" form:"problem_id" uri:"problem_id"`
	UserID    int       `json:"user_id,omitempty" form:"user_id"`
	Title     string    `json:"title,omitempty" form:"title"`
	Content   string    `json:"content,omitempty" form:"content"`
	Time      time.Time `json:"time,omitempty" form:"time" gorm:"omitempty"`
	Status    int       `json:"status,omitempty" form:"status"`
}

func (Discuss) TableName() string {
	return "discuss"
}

func (model *Discuss) GetAllDiscuss(ContestID int, ProblemID int, Offset int, Limit int) helper.ReturnType {
	var discussions []Discuss

	if ContestID != 0 {
		err := db.Offset(Offset).Limit(Limit).Where("contest_id = ?", ContestID).Find(&discussions).Error
		if err != nil {
			return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
		}
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: discussions}
	} else {
		err := db.Offset(Offset).Limit(Limit).Where("problem_id = ?", ProblemID).Find(&discussions).Error
		if err != nil {
			return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
		}
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: discussions}
	}
}

func (model *Discuss) GetDiscussionByID(id int, PageNumber int) helper.ReturnType {
	var discuss Discuss
	replyModel := Reply{}

	err := db.
		Select([]string{"id", "time", "status", "content", "title", "contest_id", "problem_id"}).
		Where("id = ?", id).
		First(&discuss).
		Error
	res := replyModel.GetReplyByDiscussID(id, (PageNumber-1)*constants.PageLimit, constants.PageLimit)

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: gin.H{
			"discuss": discuss,
			"reply":   res.Data,
		}}
	}

}

func (model *Discuss) AddDiscussion(newDiscussion Discuss) helper.ReturnType {

	err := db.Omit("time").Create(&newDiscussion).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "添加失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加成功", Data: true}
	}
}

func (model *Discuss) GetContestDiscussion(ContestID int, PageNumber int) helper.ReturnType {
	var discussions []Discuss
	var DiscussCount int
	err := db.
		Model(&Discuss{}).
		Select([]string{"id", "title", "content", "problem_id", "time"}).
		Where("contest_id = ?", ContestID).
		Count(&DiscussCount).
		Offset((PageNumber - 1) * (constants.PageDiscussLimit)).
		Limit(constants.PageDiscussLimit).
		Find(&discussions).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: gin.H{"data": discussions, "count": DiscussCount}}
	}

}

// MARK: 此接口没用
func (model *Discuss) GetProblemDiscussion(ProblemID int, PageNumber int) helper.ReturnType {
	var discussions []Discuss

	err := db.Where("problem_id = ?", ProblemID).Find(&discussions).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: discussions}
	}
}

func (model *Discuss) GetUserDiscussion(UserID int) helper.ReturnType {
	var discussions []Discuss

	err := db.
		Select([]string{"id", "title", "content", "problem_id", "time"}).
		Where("user_id = ?", UserID).
		Find(&discussions).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: discussions}
	}
}

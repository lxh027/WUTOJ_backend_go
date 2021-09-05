package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"time"
)

type Contest struct {
	ContestID   int       `json:"contest_id" form:"contest_id"`
	ContestName string    `json:"contest_name" form:"contest_name"`
	BeginTime   time.Time `json:"begin_time" form:"begin_time"`
	EndTime     time.Time `json:"end_time" form:"end_time"`
	Frozen      float64   `json:"frozen" form:"frozen"`
	Problems    string    `json:"problems" form:"problems"`
	Prize       string    `json:"prize" form:"prize"`
	Colors      string    `json:"colors" form:"colors"`
	Rule        int       `json:"rule" form:"rule"`
	Status      int       `json:"status" form:"status"`
}

func (model *Contest) GetAllContest(offset int, limit int, title string, time time.Time) helper.ReturnType {
	var contests []Contest
	where := "contest_name like ? AND begin_time > ?"
	var count int

	db.Model(&Contest{}).Where(where, "%"+title+"%", time).Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+title+"%", time).
		Find(&contests).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"contests": contests,
				"count":    count,
			},
		}
	}
}

func (model *Contest) FindContestByID(id int) helper.ReturnType {
	var contest Contest

	err := db.Where("contest_id = ?", id).First(&contest).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: contest}
	}
}

func (model *Contest) AddContest(newContest Contest) helper.ReturnType { //jun
	var contest Contest
	if err := db.Where("contest_name = ?", newContest.ContestName).First(&contest).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "比赛已存在", Data: false}
	}

	err := db.Create(&newContest).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		db.Where("contest_name = ?", newContest.ContestName).First(&contest)
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: contest.ContestID}
	}
}


func (model *Contest) DeleteContest(contestID int) helper.ReturnType {
	err := db.Where("contest_id = ?", contestID).Delete(&Contest{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Contest) UpdateContest(contestID int, updateContest Contest) helper.ReturnType {
	err := db.Model(&Contest{}).Where("contest_id = ?", contestID).Update(updateContest).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Contest) ChangeContestStatus(contestID int, status int) helper.ReturnType {
	err := db.Model(&Contest{}).Where("contest_id = ?", contestID).Update("status", status).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Contest) GetContestById(contestID uint) helper.ReturnType {
	contest := Contest{}
	err := db.Where("contest_id = ?", contestID).First(&contest).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "未找到数据", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查找成功", Data: contest}
}

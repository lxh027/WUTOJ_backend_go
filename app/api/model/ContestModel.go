package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	_ "OnlineJudge/config"
	_ "debug/elf"
	_ "github.com/go-playground/locales/mgh"
	"time"
)

type Contest struct {
	ContestID   int       `json:"contest_id" form:"contest_id" uri:"contest_id"`
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

func (Contest) TableName() string {
	return "contest"
}

func (model *Contest) GetContest(param string) helper.ReturnType {
	contest := Contest{}
	err := db.Where("contest_id = ?", param).Find(&contest).Error

	if err == nil {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: contest}
	} else {
		if err = db.Where("contest_name = ?", param).Find(&contest).Error; err == nil {
			return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: contest}
		}
	}

	return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: ""}
}

func (model *Contest) GetContestByName(contestName string) helper.ReturnType {
	contest := Contest{}
	if contestName != "" {
		err := db.Where("contest_name = ?", contestName).First(&contest).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "未找到数据", Data: err.Error()}
		} else {
			return helper.ReturnType{Status: common.CodeSuccess, Msg: "查找成功", Data: contest}
		}
	}

	err := db.Where("contest_id = ?", contestName).First(&contest).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "数据库错误", Data: ""}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查找成功", Data: contest}
	}
}

func (model *Contest) GetContestById(contestID string) helper.ReturnType {
	contest := Contest{}
	if contestID != "" {
		err := db.Where("contest_id = ?", contestID).First(&contest).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "未找到数据", Data: err.Error()}
		} else {
			return helper.ReturnType{Status: common.CodeSuccess, Msg: "查找成功", Data: contest}
		}
	}

	err := db.Where("contest_id = ?", contestID).First(&contest).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "数据库错误", Data: ""}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查找成功", Data: contest}
	}
}

func (model *Contest) GetAllContest(Offset int, Limit int) helper.ReturnType {
	var contests []Contest
	//
	//db.Model(&Contest{}).Find(&contests).Error
	//
	err := db.Offset(Offset).Limit(Limit).Find(&contests).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查找失败，数据库错误", Data: ""}
	}
	if contests != nil {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查找成功", Data: contests}
	}
	return helper.ReturnType{Status: common.CodeError, Msg: "当前无比赛", Data: ""}

}

func (model *Contest) AddContest(data Contest) helper.ReturnType {
	err := db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "新建比赛失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "新建比赛成功", Data: ""}
	}
}

func (model *Contest) DeleteContest(contestID int) helper.ReturnType {
	//contest := Contest{}
	//err := db.Where("contest_id = ?", contestID).First(&contest).Error
	//if err != nil {
	//	return helper.ReturnType{Status: common.CodeError, Msg: "数据库错误||删除比赛失败", Data: err.Error()}
	//}
	//err = db.Delete(&contest).Error
	//contest.Status = 0
	//err = db.Create(&contest).Error
	err := db.Model(model).Where("contest_id = ?", contestID).Update("status", 0).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除比赛失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除比赛成功", Data: ""}
	}
}

func (model *Contest) GetContestStatus(ContestID int) helper.ReturnType {
	var contest = Contest{}
	err := db.Where("contest_id = ?", ContestID).First(&contest).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取比赛状态失败", Data: ""}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取比赛状态成功", Data: contest.Status}
}

func (model *Contest) GetContestProblems(ContestID int) helper.ReturnType {
	var contest = Contest{}
	err := db.Where("contest_id = ?", ContestID).First(&contest).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取题目失败，数据库错误", Data: ""}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取题目成功", Data: contest.Problems}
}

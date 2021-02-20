package model

// Wait To Do

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	_ "OnlineJudge/config"
	_ "debug/elf"
	_ "github.com/go-playground/locales/mgh"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Contest struct {
	ContestID   int                  `json:"contest_id" form:"contest_id"`
	ContestName string               `json:"contest_name" form:"contest_name"`
	BeginTime   timestamp.Timestamp            `json:"begin_time form:"begin_time"`
	EndTime     timestamp.Timestamp            `json:"end_time form:"end_time"`
	Frozen      float64 `json:"frozen" form:"frozen"`
	Problems    string               `json:"problems" form:"problems"`
	Prize       string               `json:"prize" form:"prize"`
	Colors 		string				 `json:"colors" form:"colors"`
	Rule 		int 				 `json:"rule" form:"rule"`
	Status      int                  `json:"status" form:"status"`
}

func (Contest) TableName() string {
	return "contest"
}


func (model *Contest) GetContestByName(contestName string) helper.ReturnType {
	contest := Contest{}
	if contestName != "" {
		err := db.Where("contest_id = ?", contestName).First(&contest).Error
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


func (model *Contest) GetAllContest() helper.ReturnType {
	var contests []Contest
	result := db.Find(&contests)
	if result != nil {
		return helper.ReturnType{Status: common.CodeSuccess, Msg:"查找成功", Data: result.}
	} else {
		return helper.ReturnType{Status: common.CodeError, Msg: "当前无比赛", Data: ""}
	}
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
	contest := Contest{}
	err := db.Where("contest_id = ?", contestID).First(&contest).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "数据库错误||删除比赛失败", Data: err.Error()}
	}
	err = db.Model(model).Where("contest_id = ?", contestID).Update("status", 0).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除比赛失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除比赛成功", Data: ""}
	}
}

func (model *Contest) UpdateContest(contestID int, data Contest) helper.ReturnType {
	err := db.Model(model).Where("contest_id = ?", contestID).Update(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "编辑比赛失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "编辑比赛成功", Data: 1}
	}
}


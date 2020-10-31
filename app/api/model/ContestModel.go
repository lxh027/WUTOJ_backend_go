package model

import (
	"OnlineJudge/app/common"
	"debug/elf"
	"github.com/go-playground/locales/mgh"
	"github.com/golang/protobuf/ptypes/wrappers"
	"time"
)

type Contest struct {
	ContestID   int                  `json:"contest_id" form:"contest_id"`
	ContestName string               `json:"contest_name" form:"contest_name"`
	BeginTime   time.Time            `json:"begin_time form:"begin_time"`
	EndTime     time.Time            `json:"end_time form:"end_time"`
	Frozen      wrappers.DoubleValue `json:"frozen" form:"frozen"`
	Problems    string               `json:"problems" form:"problems"`
	Prize       string               `json:"prize" form:"prize"`
	GroupID     int                  `json:"group_id" form:"prize"`
	Status      int                  `json:"status" form:"status"`
}

func (Contest) TableName() string {
	return "contest"
}

func (model *Contest) searchContest(contestID int, contestName string) common.ReturnType {
	if contestID == 0 && contestName == "" {
		return common.ReturnType{Status: common.CODE_ERROE}
	}
	contest := Contest{}
	if contestID != 0 {
		err := db.Where("contest_id = ?", contestID).First(contest).Error
		if err == nil {
			return common.ReturnType{Status: common.CODE_ERROE, Msg: "未找到数据", Data: err.Error()}
		}

	}

	err := db.Where("contest_id = ?", contestID).First(contest).Error

	if err != nil {

	} else {

	}
}

func (model *Contest) newContest(data Contest) common.ReturnType {
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: data}
	}
}

func (model *Contest) deleteContest(contestID int) common.ReturnType {
	contest := Contest{}
	err := db.Where("contest_id = ?", contestID).First(contest)
	if err == nil {
		return common.ReturnType{}
	}
	err = db.Delete(&contest)
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "删除比赛失败", Data: ""}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "删除比赛成功", Data: ""}
	}
}

func (model *Contest) editContest(contestID int, data Contest) common.ReturnType {
	err := db.Model(model).Where("contest_id = ?", contestID).Update(data).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "编辑比赛失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "编辑比赛成功", Data: 1}
	}
}

func (model *Contest) getAllGroupContest(groupID int) {

}

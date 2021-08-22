package model

// //自建
// import (
// 	"OnlineJudge/app/helper"
// 	"OnlineJudge/constants"
// )

// //ContestCompetitor 竞赛与参赛者的KV表
// type ContestCompetitor struct {
// 	ContestID int `json:"contest_id" form:"contest_id"`
// 	Rid       int `json:"rid" form:"rid"`
// }

// //AddContestCompetitor 添加
// func (model *ContestCompetitor) AddContestCompetitor(newContestCompetitor ContestCompetitor) helper.ReturnType {
// 	err := db.Create(&newContestCompetitor).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
// 	}
// }

// //DeleteContestCompetitor 删除
// func (model *ContestCompetitor) DeleteContestCompetitor(newContestCompetitor ContestCompetitor) helper.ReturnType {
// 	err := db.Where("contest_id = ? AND rid = ?", newContestCompetitor.ContestID, newContestCompetitor.Rid).Delete(ContestCompetitor{}).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
// 	}
// }

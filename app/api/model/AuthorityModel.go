package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Authority struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Enabled uint8  `json:"enabled"`
}

func (model *Authority) GetAllAuthority() helper.ReturnType {
	authorities := make([]Authority, 0)

	err := db.Find(&authorities).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取成功", Data: authorities}
}

func (model *Authority) GetAuthorityByID(id uint64) helper.ReturnType {
	authority := Authority{}
	err := db.Where("id = ?", id).First(&authority).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取成功", Data: authority}
}

package model

import (
	"OnlineJudge/app/common"
	"log"
)


type Authority struct {
	ID uint64		`json:"id"`
	Name string		`json:"name"`
	Enabled uint8	`json:"enabled"`
}

func (model *Authority) GetAllAuthority() common.ReturnType  {
	authorities := make([]Authority, 0)

	db.Find(&authorities)

	if err := db.Error; err != nil {
		log.Println(err.Error())
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "获取失败", Data: err.Error()}
	}
	return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "获取成功", Data: authorities}
}

func (model *Authority) GetAuthorityByID(id uint64) common.ReturnType  {
	authority := Authority{}
	db.Where("id = ?", id).First(&authority)

	if err := db.Error; err != nil {
		log.Println(err.Error())
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "获取失败", Data: err.Error()}
	}
	return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "获取成功", Data: authority}
}



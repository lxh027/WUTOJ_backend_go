package model

import (
	"OnlineJudge/db_server"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB = db_server.MySqlDb

type Authority struct {
	ID uint64		`json:"id"`
	Name string		`json:"name"`
	Enabled uint8	`json:"enabled"`
}

func (model *Authority) GetAllAuthority() (authorities []Authority, err error)  {
	authorities = make([]Authority, 0)

	db.Find(&authorities)

	if err = db.Error; err != nil {
		log.Println(err.Error())
		return
	}
	return
}



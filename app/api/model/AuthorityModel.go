package model

import (
	"log"
)


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

func (model *Authority) GetAuthorityByID(id uint64) (authority Authority, err error)  {
	db.Where("id = ?", id).First(&authority)

	if err = db.Error; err != nil {
		log.Println(err.Error())
		return
	}
	return
}



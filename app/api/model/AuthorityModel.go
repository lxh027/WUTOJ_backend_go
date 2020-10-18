package model

import (
	"OnlineJudge/db_server"
	"database/sql"
	"log"
)

var db *sql.DB = db_server.MySqlDb

type Authority struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Enabled int64 `json:"enabled"`
}

func (model *Authority) GetAllAuthority() (authorities []Authority, err error)  {
	authorities = make([]Authority, 0)

	rows, err := db.Query("SELECT * FROM `authority`")

	if err != nil {
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		var auth Authority
		err = rows.Scan(&auth.ID, &auth.Name, &auth.Enabled)
		authorities = append(authorities, auth)
	}

	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}



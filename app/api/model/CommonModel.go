package model

import (
	"OnlineJudge/db_server"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = db_server.MySqlDb

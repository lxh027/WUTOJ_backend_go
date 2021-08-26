package model

import (
	"OnlineJudge/core/db"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = db.MySqlDb


package model

import (
	"OnlineJudge/core/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = database.MySqlDb


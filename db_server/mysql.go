package db_server

import (
	_ "github.com/go-sql-driver/mysql"

	"OnlineJudge/config"
	"database/sql"
	"fmt"
)


var MySqlDb *sql.DB;
var MySqlError error;

func init()  {
	dbConfig := config.GetDbConfig()

	// set db dsn
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			dbConfig["username"],
			dbConfig["password"],
			dbConfig["hostname"],
			dbConfig["port"],
			dbConfig["database"],
			dbConfig["charset"],
		)

	// open connection
	MySqlDb, MySqlError = sql.Open("mysql", dbDSN)

	if MySqlError != nil {
		panic("database open error! " + MySqlError.Error())
	}

	if MySqlError = MySqlDb.Ping(); MySqlError != nil {
		panic("database connect error! " + MySqlError.Error())
	}


}





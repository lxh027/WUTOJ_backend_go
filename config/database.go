package config

import "os"

var dbConfig map[string]interface{}

func init() {
	// init db config
	dbConfig = make(map[string]interface{})

	dbConfig["hostname"] = os.Getenv("database_host")
	dbConfig["port"] = os.Getenv("database_port")
	dbConfig["database"] = os.Getenv("database_name")
	dbConfig["username"] = os.Getenv("database_user")
	dbConfig["password"] = os.Getenv("database_passwd")
	dbConfig["charset"] = "utf8"
	dbConfig["parseTime"] = "True"
}

func GetDbConfig() map[string]interface{} {
	return dbConfig
}

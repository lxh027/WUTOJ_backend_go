package config

func GetDbConfig() map[string]interface{} {

	// init db config
	dbConfig := make(map[string]interface{})

	dbConfig["hostname"] 	= "localhost"
	dbConfig["port"] 		= "3306"
	dbConfig["database"] 	= "online_judge_dev"
	dbConfig["username"] 	= "root"
	dbConfig["password"] 	= "19991107"
	dbConfig["charset"]		= "utf8"
	dbConfig["parseTime"]	= "True"

	return dbConfig
}

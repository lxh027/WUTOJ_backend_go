package config

var dbConfig map[string]interface{}

func init() {
	// init db config
	dbConfig = make(map[string]interface{})

	dbConfig["hostname"] = "120.77.181.57"
	dbConfig["port"] = "3306"
	dbConfig["database"] = "online_judge_dev"
	dbConfig["username"] = "online_judge_dev"
	dbConfig["password"] = "12345678"
	dbConfig["charset"] = "utf8"
	dbConfig["parseTime"] = "True"
}

func GetDbConfig() map[string]interface{} {
	return dbConfig
}

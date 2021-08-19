package config

var dbConfig map[string]interface{}

func init() {
	// init db config
	dbConfig = make(map[string]interface{})

	// dbConfig["hostname"] = os.Getenv("database_host")
	dbConfig["hostname"] = "localhost"
	// dbConfig["port"] = os.Getenv("database_port")
	dbConfig["port"] = "3306"
	// dbConfig["database"] = os.Getenv("database_name")
	dbConfig["database"] = "online_judge"
	// dbConfig["username"] = os.Getenv("database_user")
	dbConfig["username"] = "root"
	// dbConfig["password"] = os.Getenv("database_passwd")
	dbConfig["password"] = "root"
	dbConfig["charset"] = "utf8"
	dbConfig["parseTime"] = "True"

	dbConfig["maxIdleConns"] = 20
	dbConfig["maxOpenConns"] = 100
	/*dbConfig["connMaxIdleTime"] = 10000
	dbConfig["connMaxLifetime"] = 600000*/

}

func GetDbConfig() map[string]interface{} {
	return dbConfig
}

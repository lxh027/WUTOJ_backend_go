package config

var sessionConfig map[string]interface{}

func init() {
	sessionConfig = make(map[string]interface{})

	sessionConfig["key"] = "online_judge"
	sessionConfig["name"] = "oj_session"
	sessionConfig["age"] = 86400
	sessionConfig["path"] = "/"
}

func GetSessionConfig() map[string]interface{} {
	return sessionConfig
}

package config

func GetSessionConfig() map[string]interface{}{
	sessionConfig := make(map[string]interface{})

	sessionConfig["key"] 	= "online_judge"
	sessionConfig["name"]	= "oj_session"
	sessionConfig["age"] 	= 86400
	sessionConfig["path"]	= "/"
	return sessionConfig
}
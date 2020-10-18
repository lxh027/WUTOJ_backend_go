package config

func GetSessionConfig() map[string]string{
	sessionConfig := make(map[string]string)

	sessionConfig["key"] 	= "online_judge"
	sessionConfig["name"]	= "oj_session"

	return sessionConfig
}
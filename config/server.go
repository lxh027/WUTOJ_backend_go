package config

func GetServerConfig() map[string]string {
	serverConfig := make(map[string]string)

	serverConfig["host"] 	= "0.0.0.0"
	serverConfig["port"] 	= "5000"

	serverConfig["mode"]	= "debug"
	return serverConfig
}

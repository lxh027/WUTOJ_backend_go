package config

var serverConfig map[string]interface{}

func init() {
	serverConfig = make(map[string]interface{})

	serverConfig["host"] = "0.0.0.0"
	serverConfig["port"] = "5000"

	serverConfig["mode"] = "debug"
}

func GetServerConfig() map[string]interface{} {
	return serverConfig
}

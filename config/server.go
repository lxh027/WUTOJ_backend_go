package config

import "os"

var serverConfig map[string]interface{}

func init() {
	serverConfig = make(map[string]interface{})

	serverConfig["host"] = "0.0.0.0"
	serverConfig["port"] = os.Getenv("server_port")
	serverConfig["domain"] = "acmwhut.com"
	serverConfig["mode"] = "debug"
}

func GetServerConfig() map[string]interface{} {
	return serverConfig
}

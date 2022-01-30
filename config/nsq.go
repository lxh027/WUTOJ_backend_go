package config

var nsqConfig map[string]interface{}

func init() {
	// init db config
	nsqConfig = make(map[string]interface{})

	nsqConfig["host"] = "localhost"
	nsqConfig["port"] = "4150"
	nsqConfig["topic"] = "crawler"

}

func GetNSQConfig() map[string]interface{} {
	return nsqConfig
}

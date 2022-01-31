package config

import "os"

var nsqConfig map[string]interface{}

func init() {
	// init db config
	nsqConfig = make(map[string]interface{})

	nsqConfig["host"] = os.Getenv("nsq_host")
	nsqConfig["port"] = os.Getenv("nsq_port")
	nsqConfig["producer_topic"] = os.Getenv("nsq_producer_topic")
	nsqConfig["consumer_topic"] = os.Getenv("nsq_consumer_topic")
	// nsqConfig["host"] = "localhost"
	// nsqConfig["port"] = "4150"
	// nsqConfig["producer_topic"] = "crawler"
	// nsqConfig["consumer_topic"] = "responce"

}

func GetNSQConfig() map[string]interface{} {
	return nsqConfig
}

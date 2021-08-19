package config

import (
	"os"
	"time"
)

var redisConfig map[string]interface{}

func init() {
	redisConfig = make(map[string]interface{})

	redisConfig["env"] = os.Getenv("env")
	// redisConfig["env"] = "dev"
	redisConfig["rank_cache_time"] = 5
	redisConfig["host"] = os.Getenv("redis_host")
	redisConfig["auth"] = ""
	redisConfig["type"] = "tcp"

	// 初始连接数量
	redisConfig["maxIdle"] = 16
	// 最大连接数量
	redisConfig["maxActive"] = 0
	// 过期时间
	redisConfig["timeout"] = 300 * time.Second

}

func GetRedisConfig() map[string]interface{} {
	return redisConfig
}

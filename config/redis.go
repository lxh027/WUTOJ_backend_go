package config

import "time"

func GetRedisConfig()map[string]interface{}  {
	redisConfig := make(map[string]interface{})

	redisConfig["rank_cache_time"] 	= 5
	redisConfig["host"]				= "localhost:6379"
	redisConfig["auth"]				= ""
	redisConfig["type"]				= "tcp"

	// 初始连接数量
	redisConfig["maxIdle"]			= 16
	// 最大连接数量
	redisConfig["maxActive"]		= 0
	// 过期时间
	redisConfig["timeout"]			= 300 * time.Second

	return redisConfig
}
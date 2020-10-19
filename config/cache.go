package config

func GetCacheConfig()map[string]interface{}  {
	cacheConfig := make(map[string]interface{})

	cacheConfig["rank_cache_time"] 	= 5
	cacheConfig["host"]				= "localhost"
	cacheConfig["port"]				= "6379"

	return cacheConfig
}
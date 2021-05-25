package config

var cacheConfig map[string]interface{}

func init() {
	cacheConfig = make(map[string]interface{})

	cacheConfig["rank_cache_time"] = 5
	cacheConfig["host"] = "localhost"
	cacheConfig["port"] = "6379"
}

func GetCacheConfig() map[string]interface{} {
	return cacheConfig
}

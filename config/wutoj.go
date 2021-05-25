package config

func GetWutOjConfig() map[string]interface{} {
	wutOjConfig := make(map[string]interface{})

	/* 排行榜存储时间 */
	wutOjConfig["rank_cache_time"] = 5
	wutOjConfig["host"] = "localhost"
	wutOjConfig["port"] = "6379"
	wutOjConfig["user_rank_cache"] = "user_rank_cache"
	wutOjConfig["environment"] = "master"

	/* 提交 */
	wutOjConfig["submit_url"] = []string{
		"http://10.143.216.128:8819/submit",
	}

	/* 交题时间间隔 */
	wutOjConfig["interval_time"] = 0

	/* 打印请求间隔 */
	wutOjConfig["print_interval_time"] = 0

	/* 支持语言 */
	wutOjConfig["language"] = []string{
		"c.gcc",
		"cpp.g++",
		"java.openjdk",
		"py.cpython",
	}

	/* 每面数量 */
	wutOjConfig["page_limit"] = 20

	/* oj链接 */
	wutOjConfig["oj_url"] = "http://acmwhut.com"
	return wutOjConfig
}

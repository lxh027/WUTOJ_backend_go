package config

func GetJudgeConfig()map[string]interface{}  {
	judgeConfig := make(map[string]interface{})

	judgeConfig["env"] = "dev"
	judgeConfig["address"] = "127.0.0.1:8800"
	judgeConfig["base_dir"] = "/home/acmwhut/data"
	judgeConfig["tmp_dir"] = "/tmpdir"

	return judgeConfig
}
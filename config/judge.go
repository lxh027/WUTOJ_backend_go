package config

func GetJudgeConfig() map[string]interface{} {
	judgeConfig := make(map[string]interface{})

	judgeConfig["env"] = ""
	judgeConfig["address"] = "127.0.0.1:8800"
	judgeConfig["base_dir"] = "/home/acmwhut/data/dev"
	judgeConfig["tmp_dir"] = "/home/ana_tmpdir"

	return judgeConfig
}

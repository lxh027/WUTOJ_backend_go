package config

type LangConfig struct {
	Lang         string
	BuildSh      string
	RunnerConfig string
}

func GetJudgeConfig() map[string]interface{} {
	judgeConfig := make(map[string]interface{})

	judgeConfig["env"] = "dev"
	judgeConfig["address"] = "127.0.0.1:8800"
	judgeConfig["base_dir"] = "/home/acmwhut/data"
	judgeConfig["tmp_dir"] = "/home/ana_tmpdir"

	return judgeConfig
}

func GetLangConfigs() []LangConfig {
	langBasePath := "/home/baka233/acmwhut/env"

	langConfigs := []LangConfig{
		{"c.gcc", "", langBasePath + "/c.gcc/runner.toml"},
		{"py.cpython3.6", langBasePath + "/py.cpython3.6/build.sh", langBasePath + "/py.cpython3.6/runner.toml"},
		{"java.openjdk-10", langBasePath + "/java.openjdk-10/build.sh", langBasePath + "/java.openjdk-10/runner.toml"},
	}
	return langConfigs
}

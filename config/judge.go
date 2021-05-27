package config

import "os"

type LangConfig struct {
	Lang         string
	BuildSh      string
	RunnerConfig string
}

var judgeConfig map[string]interface{}
var langConfigs []LangConfig

func init() {
	judgeConfig = make(map[string]interface{})

	// 初始化 judge config

	judgeConfig["env"] = "dev"
	judgeConfig["address"] = os.Getenv("ana_addr")
	judgeConfig["base_dir"] = os.Getenv("data")
	judgeConfig["tmp_dir"] = "/home/ana_tmpdir"

	// 初始化 lang config

	langBasePath := "/home/env"

	langBuildPath := []string{
		"",
		"",
		"/java.openjdk8/build.sh",
		"/python.cpython3.6/build.sh",
	}

	langRunnerConfig := []string{
		"/c.gcc/runner.toml",
		"/cpp.g++/runner.toml",
		"/java.openjdk8/runner.toml",
		"/python.cpython3.6/runner.toml",
	}

	langConfigs = []LangConfig{
		{"c.gcc", "", langBasePath + langRunnerConfig[0]},
		{"cpp.g++", "", langBasePath + langRunnerConfig[1]},
		{"java.openjdk8", langBasePath + langBuildPath[2], langBasePath + langRunnerConfig[2]},
		{"python.cpython3.6", langBasePath + langBuildPath[3], langBasePath + langRunnerConfig[3]},
	}

}

func GetJudgeConfig() map[string]interface{} {
	return judgeConfig
}

func GetLangConfigs() []LangConfig {
	return langConfigs
}

package config

func GetMailConfig() map[string]interface{} {
	mailConfig := make(map[string]interface{})

	mailConfig["charset"] = "utf-8"
	mailConfig["smtp_debug"] = 0
	mailConfig["host"] = "smtp.163.com"
	mailConfig["smtp_secure"] = "ssl"
	mailConfig["port"] = 465
	mailConfig["username"] = "acmwut@163.com"
	mailConfig["password"] = "acmwut123"
	mailConfig["from"] = "acmwut@163.com"
	mailConfig["from_name"] = "武汉理工大学ACM协会"
	mailConfig["address"] = "127.0.0.1:8800"

	return mailConfig
}

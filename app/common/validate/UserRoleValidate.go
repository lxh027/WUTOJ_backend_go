package validate

import "OnlineJudge/app/helper"

var UserRoleValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id":  "required",
		"rid":  "required",
		"rids": "required",
	}

	scenes := map[string][]string{
		"add":         {"user_id", "rid"},
		"addGroup":    {"user_id", "rids"},
		"deleteGroup": {"user_id", "rids"},
		"delete":      {"user_id", "rid"},
		"getUserRole": {"user_id"},
	}

	UserRoleValidate.Rules = rules
	UserRoleValidate.Scenes = scenes
}

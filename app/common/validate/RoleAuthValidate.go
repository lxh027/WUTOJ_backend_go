package validate

import "OnlineJudge/app/helper"

var RoleAuthValidate helper.Validator

func init() {
	rules := map[string]string {
		"rid"	: "required",
		"aid"	: "required",
		"aids" 	: "required",
	}

	scenes := map[string] []string {
		"add" : {"rid", "aid"},
		"addGroup": {"rid", "aids"},
		"deleteGroup": {"rid", "aids"},
		"delete": {"rid", "aid"},
		"getRoleAuth": {"rid"},
	}

	RoleAuthValidate.Rules = rules
	RoleAuthValidate.Scenes = scenes
}

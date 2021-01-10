package validate

import "OnlineJudge/app/helper"

var RoleValidate helper.Validator

func init() {
	rules := map[string]string {
		"rid"	: "required",
		"name"	: "required",
		"desc"	: "required",
	}

	scenes := map[string] []string {
		"add" : {"name", "desc"},
		"delete": {"rid"},
		"find"	: {"rid"},
		"update": {"rid", "desc", "name"},
	}

	RoleValidate.Rules = rules
	RoleValidate.Scenes = scenes
}

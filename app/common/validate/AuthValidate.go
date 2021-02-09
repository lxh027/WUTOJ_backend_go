package validate

import "OnlineJudge/app/helper"

var AuthValidate helper.Validator

func init() {
	rules := map[string]string {
		"aid"	: "required",
		"icon"	: "required",
		"title"	: "required",
		"type"	: "required",
	}

	scenes := map[string] []string {
		"add" : {"title", "type", "icon"},
		"delete": {"aid"},
		"find"	: {"aid"},
		"findParent": {"type"},
		"update": {"aid", "title", "icon"},
	}

	AuthValidate.Rules = rules
	AuthValidate.Scenes = scenes
}
